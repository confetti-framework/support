package value

import (
	"github.com/lanvard/support"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFloatFromEmptyString(t *testing.T) {
	value := support.NewValue("")

	result, err := value.FloatE()

	assert.Equal(t, 0.0, result)
	assert.Error(t, err, "unable to cast \"\" of type string to float64")
}

func TestFloatFromWords(t *testing.T) {
	value := support.NewValue("four")

	result, err := value.FloatE()

	assert.Equal(t, 0.0, result)
	assert.EqualError(t, err, "unable to cast \"four\" of type string to float64")
}

func TestFloatFromLongNumber(t *testing.T) {
	value := support.NewValue(
		"12345678912345367891234523456567896123475123456789123453678912345234565678" +
			"9612347567912345678912345367891234567912344567891253456789612347567912" +
			"3456789123453678912345679123445678912534567896123475679123456789123453" +
			"6789123456791234456789125345678961234756791234567891234536789123456791" +
			"23445678912534567896123475679",
	)

	result, err := value.FloatE()

	assert.Equal(t, 0.0, result)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "unable to cast \"1234")
	assert.Contains(t, err.Error(), "475679\" of type string to float64")
}

func TestFloatFromString(t *testing.T) {
	value := support.NewValue("1.5123")

	result, err := value.FloatE()

	assert.Equal(t, 1.5123, result)
	assert.NoError(t, err)
}

func TestFloatFromDifferentIntTypes(t *testing.T) {
	var result float64

	result, _ = support.NewValue(312).FloatE()
	assert.Equal(t, float64(312), result)

	result, _ = support.NewValue(int8(2)).FloatE()
	assert.Equal(t, float64(2), result)

	result, _ = support.NewValue(int16(2)).FloatE()
	assert.Equal(t, float64(2), result)
}

func TestFirstFloatFromCollection(t *testing.T) {
	result := support.NewValue(support.NewCollection(12.12)).Float()

	assert.Equal(t, 12.12, result)
}

func TestFirstFloatFromMap(t *testing.T) {
	input := support.NewMap(map[string]interface{}{"total": 12.12})
	result := support.NewValue(input).Float()

	assert.Equal(t, 12.12, result)
}

func TestFloatNotPanicWithoutErrorReceiver(t *testing.T) {
	assert.NotPanics(t, func() {
		support.NewValueE(123, nil).Float()
	})
	assert.Equal(t, float64(123), support.NewValueE(123, nil).Float())
}

func TestFloatPanicWithoutErrorReceiver(t *testing.T) {
	assert.PanicsWithError(t, "error_message", func() {
		support.NewValueE("test", "error_message").Float()
	})
}
