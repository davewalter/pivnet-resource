package acceptance

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"

	"github.com/onsi/gomega/gexec"
	"github.com/pivotal-cf/go-pivnet/v4"
	"github.com/pivotal-cf/go-pivnet/v4/logshim"
	"github.com/pivotal-cf/pivnet-resource/gp"
	"github.com/robdimsdale/sanitizer"

	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	inPath    string
	checkPath string
	outPath   string

	endpoint string

	productSlug  string
	refreshToken string

	pivnetClient                      *gp.Client
	additionalSynchronizedBeforeSuite func(SuiteEnv)
)

func TestAcceptance(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Acceptance Suite")
}

type SuiteEnv struct {
	InPath    string
	CheckPath string
	OutPath   string

	Endpoint string

	ProductSlug  string
	RefreshToken string
}

var _ = SynchronizedBeforeSuite(func() []byte {
	suiteEnv := SuiteEnv{}
	var err error
	By("Getting product slug from environment variables")
	suiteEnv.ProductSlug = os.Getenv("PRODUCT_SLUG")
	Expect(suiteEnv.ProductSlug).NotTo(BeEmpty(), "$PRODUCT_SLUG must be provided")

	By("Getting endpoint from environment variables")
	suiteEnv.Endpoint = os.Getenv("PIVNET_ENDPOINT")
	Expect(suiteEnv.Endpoint).NotTo(BeEmpty(), "$PIVNET_ENDPOINT must be provided")

	By("Getting refresh token from environment variables")
	suiteEnv.RefreshToken = os.Getenv("PIVNET_RESOURCE_REFRESH_TOKEN")
	Expect(suiteEnv.RefreshToken).NotTo(BeEmpty(), "$PIVNET_RESOURCE_REFRESH_TOKEN must be provided")

	By("Compiling check binary")
	suiteEnv.CheckPath, err = gexec.Build("github.com/pivotal-cf/pivnet-resource/cmd/check", "-race")
	Expect(err).NotTo(HaveOccurred())

	By("Compiling out binary")
	suiteEnv.OutPath, err = gexec.Build("github.com/pivotal-cf/pivnet-resource/cmd/out", "-race")
	Expect(err).NotTo(HaveOccurred())

	By("Compiling in binary")
	suiteEnv.InPath, err = gexec.Build("github.com/pivotal-cf/pivnet-resource/cmd/in")
	Expect(err).NotTo(HaveOccurred())

	By("Sanitizing suite setup output")
	ls := getLogShim()

	By("Creating pivnet client for suite setup")
	pivnetClient = getClient(ls, suiteEnv.Endpoint, suiteEnv.RefreshToken)

	if additionalSynchronizedBeforeSuite != nil {
		additionalSynchronizedBeforeSuite(suiteEnv)
	}

	envBytes, err := yaml.Marshal(suiteEnv)
	Expect(err).ShouldNot(HaveOccurred())
	return envBytes
}, func(envBytes []byte) {
	suiteEnv := SuiteEnv{}
	err := yaml.Unmarshal(envBytes, &suiteEnv)
	Expect(err).ShouldNot(HaveOccurred())

	inPath = suiteEnv.InPath
	checkPath = suiteEnv.CheckPath
	outPath = suiteEnv.OutPath
	endpoint = suiteEnv.Endpoint
	productSlug = suiteEnv.ProductSlug
	refreshToken = suiteEnv.RefreshToken

	By("Sanitizing acceptance test output")
	ls := getLogShim()

	By("Creating pivnet client (for out-of-band operations)")
	pivnetClient = getClient(ls, suiteEnv.Endpoint, suiteEnv.RefreshToken)
})

func getLogShim() *logshim.LogShim {
	sanitized := map[string]string{
		refreshToken: "***sanitized-refresh-token***",
	}
	sanitizedWriter := sanitizer.NewSanitizer(sanitized, GinkgoWriter)
	GinkgoWriter = sanitizedWriter

	testLogger := log.New(sanitizedWriter, "", log.LstdFlags)
	verbose := true
	return logshim.NewLogShim(testLogger, testLogger, verbose)
}

func getClient(ls *logshim.LogShim, endpoint string, refreshToken string) *gp.Client {
	clientConfig := pivnet.ClientConfig{
		Host:      endpoint,
		UserAgent: "pivnet-resource/integration-test",
	}

	return gp.NewClient(pivnet.NewAccessTokenOrLegacyToken(refreshToken, endpoint, false), clientConfig, ls)
}

var _ = SynchronizedAfterSuite(func() {
}, func() {
	gexec.CleanupBuildArtifacts()
})
