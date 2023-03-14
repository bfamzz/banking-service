package mail

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/bfamzz/banking-service/util"
	"github.com/stretchr/testify/require"
)

func TestSendEmailWithAwsSES(t *testing.T) {
	config, err := util.LoadConfig("..")
	require.NoError(t, err)
	require.NotEmpty(t, config)

	sdkConfig, err := util.LoadAwsSdkConfig()
	require.NoError(t, err)
	require.NotEmpty(t, sdkConfig)

	fromEmailNameAndAddress := fmt.Sprintf("%s <%s>", config.EmailSenderName, config.EmailSenderAddress)
	sender := NewSesSender(sdkConfig, config.EmailSenderAddress, fromEmailNameAndAddress)
	templateData := map[string]string{
		"website":"https://www.famzzie.com",
	}
	templateDataString, err := json.Marshal(templateData)
	require.NoError(t, err)
	require.NotEmpty(t, templateDataString)
	err = sender.SendTemplateEmail(string(templateDataString), []string{ "testrasr@gmail.com" }, nil, nil, nil)
	require.NoError(t, err)
}
