package httpfs

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// A https File implements a subset of os.File's methods.
type File struct {
	name   string // local file name passed to Open
	client *Client
	u      url.URL // url to access file on remote machine
	fd     uintptr // file descriptor on server
}

func (f *File) Read(p []byte) (n int, err error) {
	// send OPEN request
	u := f.u // copy
	q := u.Query()
	q.Set("n", fmt.Sprint(len(p))) // number of bytes to read
	u.RawQuery = q.Encode()
	req, eReq := http.NewRequest("READ", u.String(), nil)
	panicOn(eReq)
	resp, eResp := f.client.client.Do(req)
	if eResp != nil {
		return 0, fmt.Errorf(`httpfs read "%v": %v`, f.name, eResp)
	}
	defer resp.Body.Close()
	nRead, eRead := resp.Body.Read(p)

	errStr := resp.Header.Get(X_READ_ERROR)
	if errStr != "" {
		return nRead, errors.New(errStr) // other than EOF error
	}
	return nRead, eRead // passes on EOF
}

func (f *File) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (f *File) Close() error {
	return nil
}
