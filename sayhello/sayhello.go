package sayhello

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"httpHello/ports"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	html := `<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8">
	<title>httpHello</title>
  </head>
  <body>
    <p>
       Hello!<br>
          %s

    </p>
  </body>
</html>
`
	url := r.URL.Path
	host := r.Host
	datetime := t.Format("2006-01-02 15:04:05")
	msg := fmt.Sprintf("\nTime is %s and you are calling from %s", datetime, host)

	switch url {
	case "/404":
		http.Error(w, "404 Not Found "+msg, http.StatusNotFound)
	case "/403":
		http.Error(w, "403 Forbidden "+msg, http.StatusForbidden)
	case "/408":
		http.Error(w, "408 Request Timeout "+msg, http.StatusRequestTimeout)
	case "/410":
		http.Error(w, "410 Gone "+msg, http.StatusGone)
	case "/500":
		http.Error(w, "500 Internal Server Error "+msg, http.StatusInternalServerError)
	case "/503":
		http.Error(w, "503 Service Unavailable "+msg, http.StatusServiceUnavailable)
	default:
		_, _ = fmt.Fprintf(w, html, msg)
	} // end switch

}

func StartListener(port string) error {

	// Is it an integer?
	n, err := strconv.Atoi(port)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "<%s> is not a valid port to listen on: %s", port, err)
		return err
	}

	// in valid port range?
	if n > 65535 || n < 1024 {
		_, _ = fmt.Fprintf(os.Stderr, "<%s> is not a valid portnumber", port)
		return err
	}

	if ports.Available(port) {
		http.HandleFunc("/", sayHello)

		err := http.ListenAndServe(":"+port, nil)

		if err != nil {
			return err
		}

	} else {
		e := fmt.Sprintf("Couldn't listen on port <%q>", port)
		_, _ = fmt.Fprintf(os.Stderr, e)
		return errors.New(e)
	}

	// no problems
	return nil
}