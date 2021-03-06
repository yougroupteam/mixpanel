package mixpanel

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestClient_Track(t *testing.T) {
	viper.SetEnvPrefix("mixpanel")
	viper.AutomaticEnv()
	viper.AllowEmptyEnv(true)
	token := viper.Get("token")
	if token == nil {
		t.Fatalf("Mixpanel Token is not provided for the test. You can add MIXPANEL_TOKEN to your environment variable for the test.")
	}
	client := NewClient(token.(string))

	props := map[string]interface{}{"test": "testing"}
	event := NewEvent("go-test", props)
	event.DistinctID = "1"
	result, err := client.Track(event)
	assert.NoError(t, err)
	assert.True(t, result)

	richEvent := NewEvent("go-test", props)
	richEvent.DistinctID = "2"
	richEvent.Time = uint(time.Now().Unix())
	richEvent.IP = "8.8.8.8"
	richEvent.GroupKey = "MPGO"
	richEvent.GroupID = "MPGOTEST"
	result, err = client.Track(richEvent)
	assert.NoError(t, err)
	assert.True(t, result)
}

func TestClient_Update(t *testing.T) {
	viper.SetEnvPrefix("mixpanel")
	viper.AutomaticEnv()
	viper.AllowEmptyEnv(true)
	token := viper.Get("token")
	if token == nil {
		t.Fatalf("Mixpanel Token is not provided for the test. You can add MIXPANEL_TOKEN to your environment variable for the test.")
	}
	client := NewClient(token.(string))

	distinctID := "1"
	props := map[string]interface{}{"test": "testing"}
	update := NewSetOperation(distinctID, props)
	result, err := client.Update(update)
	assert.NoError(t, err)
	assert.True(t, result)
}

func TestUpdateOperation_JSON(t *testing.T) {
	operation := NewSetOperation("1", map[string]interface{}{"test": "testing"})
	operation.SetToken("dummy-token")
	operation.(*SetOperation).IgnoreAlias = true
	operation.(*SetOperation).IgnoreTime = false

	assert.EqualValues(t, "{\"$token\":\"dummy-token\",\"$distinct_id\":\"1\",\"$ignore_alias\":true,\"$set\":{\"test\":\"testing\"}}", UpdateOperation(operation).JSON())
}

func TestSetOperation_JSON(t *testing.T) {
	operation := NewSetOperation("1", map[string]interface{}{"test": "testing"})
	operation.SetToken("dummy-token")
	operation.(*SetOperation).IgnoreAlias = true
	operation.(*SetOperation).IgnoreTime = false
	assert.EqualValues(t, "{\"$token\":\"dummy-token\",\"$distinct_id\":\"1\",\"$ignore_alias\":true,\"$set\":{\"test\":\"testing\"}}", operation.JSON())
}

func TestNewSetOnceOperation_JSON(t *testing.T) {
	operation := NewSetOnceOperation("1", map[string]interface{}{"test": "testing"})
	operation.SetToken("dummy-token")
	assert.EqualValues(t, "{\"$token\":\"dummy-token\",\"$distinct_id\":\"1\",\"$set_once\":{\"test\":\"testing\"}}", operation.JSON())
}

func TestNewAddOperation_JSON(t *testing.T) {
	operation := NewAddOperation("1", map[string]interface{}{"test": "testing"})
	operation.SetToken("dummy-token")
	assert.EqualValues(t, "{\"$token\":\"dummy-token\",\"$distinct_id\":\"1\",\"$add\":{\"test\":\"testing\"}}", operation.JSON())
}

func TestNewAppendOperation_JSON(t *testing.T) {
	operation := NewAppendOperation("1", map[string]interface{}{"powertest": "power!!!"})
	operation.SetToken("dummy-token")
	assert.EqualValues(t, "{\"$token\":\"dummy-token\",\"$distinct_id\":\"1\",\"$append\":{\"powertest\":\"power!!!\"}}", operation.JSON())
}

func TestNewUnsetOperation(t *testing.T) {
	operation := NewUnsetOperation("1", []string{"test", "testing"})
	operation.SetToken("dummy-token")
	assert.EqualValues(t, "{\"$token\":\"dummy-token\",\"$distinct_id\":\"1\",\"$unset\":[\"test\",\"testing\"]}", operation.JSON())
}

func TestNewUnionOperation(t *testing.T) {
	operation := NewUnionOperation("1", map[string][]interface{}{"test": {"testing", "more testing"}})
	operation.SetToken("dummy-token")
	assert.EqualValues(t, "{\"$token\":\"dummy-token\",\"$distinct_id\":\"1\",\"$union\":{\"test\":[\"testing\",\"more testing\"]}}", operation.JSON())
}

func TestRemovalOperation_JSON(t *testing.T) {
	operation := NewRemovalOperation("1", map[string]interface{}{"test": "testing"})
	operation.SetToken("dummy-token")
	assert.EqualValues(t, "{\"$token\":\"dummy-token\",\"$distinct_id\":\"1\",\"$remove\":{\"test\":\"testing\"}}", operation.JSON())
}

func TestNewDeleteOperation_JSON(t *testing.T) {
	operation := NewDeleteOperation("1")
	operation.SetToken("dummy-token")
	assert.EqualValues(t, "{\"$token\":\"dummy-token\",\"$distinct_id\":\"1\",\"$delete\":\"\"}", operation.JSON())
}
