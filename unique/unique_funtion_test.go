package unique

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const in = `I love music.
I love music.
I love music.

I love music of Kartik.
I love music of Kartik.
Thanks.
I love music of Kartik.
I love music of Kartik.`

func TestUniqDefault(t *testing.T) {
	message := "Test default work function without parametr's"

	opts := Options{
		Count:    false,
		Double:   false,
		Unique:   false,
		Fields:   0,
		Strings:  0,
		Ignorant: false,
	}

	out := `I love music.

I love music of Kartik.
Thanks.
I love music of Kartik.`

	inputSlice := strings.Split(in, "\n")
	expected := strings.Split(out, "\n")

	actual, _ := Unique(inputSlice, opts)
	require.Equal(t, expected, actual, message)
}

func TestUniqFlagC(t *testing.T) {
	message := "Test function with flag -c "

	opts := Options{
		Count:    true,
		Double:   false,
		Unique:   false,
		Fields:   0,
		Strings:  0,
		Ignorant: false,
	}

	out := `3 I love music.
1 
2 I love music of Kartik.
1 Thanks.
2 I love music of Kartik.`

	inputSlice := strings.Split(in, "\n")
	expected := strings.Split(out, "\n")

	actual, _ := Unique(inputSlice, opts)

	require.Equal(t, expected, actual, message)
}

func TestUniqFlagD(t *testing.T) {
	message := "Test function with flag -d"

	opts := Options{
		Count:    false,
		Double:   true,
		Unique:   false,
		Fields:   0,
		Strings:  0,
		Ignorant: false,
	}

	out := `I love music.
I love music of Kartik.
I love music of Kartik.`

	inputSlice := strings.Split(in, "\n")
	expected := strings.Split(out, "\n")

	actual, _ := Unique(inputSlice, opts)

	require.Equal(t, expected, actual, message)
}

func TestUniqFlagU(t *testing.T) {
	message := "Test function with flag -u"

	opts := Options{
		Count:    false,
		Double:   false,
		Unique:   true,
		Fields:   0,
		Strings:  0,
		Ignorant: false,
	}

	out := `
Thanks.`

	inputSlice := strings.Split(in, "\n")
	expected := strings.Split(out, "\n")

	actual, _ := Unique(inputSlice, opts)

	require.Equal(t, expected, actual, message)
}

func TestUniqFlagI(t *testing.T) {
	message := "Test function with flag -i"

	opts := Options{
		Count:    false,
		Double:   false,
		Unique:   false,
		Fields:   0,
		Strings:  0,
		Ignorant: true,
	}

	inFlagI := `I LOVE MUSIC.
I love music.
I LoVe MuSiC.

I love MuSIC of Kartik.
I love music of kartik.
Thanks.
I love music of kartik.
I love MuSIC of Kartik.`

	out := `I LOVE MUSIC.

I love MuSIC of Kartik.
Thanks.
I love music of kartik.`

	inputSlice := strings.Split(inFlagI, "\n")
	expected := strings.Split(out, "\n")

	actual, _ := Unique(inputSlice, opts)

	require.Equal(t, expected, actual, message)
}

func TestUniqFlagF(t *testing.T) {
	message := "Test function with flag -f"

	opts := Options{
		Count:    false,
		Double:   false,
		Unique:   false,
		Fields:   1,
		Strings:  0,
		Ignorant: false,
	}

	inFlagF := `We love music.
I love music.
They love music.

I love music of Kartik.
We love music of Kartik.
Thanks.`

	out := `We love music.

I love music of Kartik.
Thanks.`

	inputSlice := strings.Split(inFlagF, "\n")
	expected := strings.Split(out, "\n")

	actual, _ := Unique(inputSlice, opts)

	require.Equal(t, expected, actual, message)
}

func TestUniqFlagS(t *testing.T) {
	message := "Test function with flag -s"

	opts := Options{
		Count:    false,
		Double:   false,
		Unique:   false,
		Fields:   0,
		Strings:  1,
		Ignorant: false,
	}

	inFlagS := `I love music.
A love music.
C love music.

I love music of Kartik.
We love music of Kartik.
Thanks.`

	out := `I love music.

I love music of Kartik.
We love music of Kartik.
Thanks.`

	inputSlice := strings.Split(inFlagS, "\n")
	expected := strings.Split(out, "\n")

	actual, _ := Unique(inputSlice, opts)

	require.Equal(t, expected, actual, message)
}

func TestUniqFalseFlag(t *testing.T) {
	message := "Test function with false flags, should return an error"

	opts := Options{
		Count:    true,
		Double:   true,
		Unique:   true,
		Fields:   0,
		Strings:  0,
		Ignorant: false,
	}

	inputSlice := []string{}

	_, err := Unique(inputSlice, opts)

	require.Error(t, err, message)
}
