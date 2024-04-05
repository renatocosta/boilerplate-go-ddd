package logfile

import (
	"testing"
)

func TestShouldFailToFileWhileReading(t *testing.T) {

	/*controll := gomock.NewController(t)

	repo := mock.NewMockLogFileRepository(controll)

	repo.EXPECT().ReadFrom(gomock.Any()).Return(nil, ErrFileNotFound)

	var tests = []struct {
		path string
		want error
	}{
		{"", ErrNameEmpty},
		{"/usr/xx/stts.csv", ErrFileNotFound},
		//{noContentPathFile, ErrFileContentSize},
	}

	for _, test := range tests {
		lfile := NewLogFile(uuid.New(), "")
		_, err := lfile.ReadFrom(test.path)
		assert.Equal(t, err, test.want, "Expected: %d - Got: %d", test.want, err)
	}*/

}
