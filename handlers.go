package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

var input Attendee

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	response := `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Attendance Form</title>

<style>
    body {
        font-family: "Inter", sans-serif;
        background: linear-gradient(135deg, #eef2f3, #dfe9f3);
        padding: 20px;
        margin: 0;
        display: flex;
        justify-content: center;
        align-items: center;
        min-height: 100vh;
    }

    .form-container {
        background: white;
        width: 100%;
        max-width: 450px;
        padding: 35px;
        border-radius: 18px;
        box-shadow: 0 8px 22px rgba(0,0,0,0.12);
        animation: fadeIn 0.5s ease;
    }

    @keyframes fadeIn {
        from {opacity: 0; transform: translateY(10px);}
        to {opacity: 1; transform: translateY(0);}
    }

    h2 {
        text-align: center;
        margin-bottom: 30px;
        color: #333;
        font-size: 22px;
    }

    /* ---- PARALLEL INPUT STYLING ---- */
    .input-group {
        margin-bottom: 25px;
        position: relative;
    }

    label {
        font-size: 14px;
        color: #666;
        display: block;
        margin-bottom: 4px;
        font-weight: 500;
    }

    input, select {
        width: 100%;
        border: none;
        border-bottom: 2px solid #ccc;
        padding: 12px 2px;
        background: transparent;
        font-size: 16px;
        transition: border-color 0.25s, box-shadow 0.25s;
    }

    input:focus, select:focus {
        outline: none;
        border-bottom: 2px solid #8ebe20;
        box-shadow: 0 6px 8px -6px rgba(106, 141, 245, 0.5);
    }

    /* Dropdown arrow style */
    select {
        appearance: none;
        background-image: url("data:image/svg+xml;utf8,<svg fill='gray' height='20' viewBox='0 0 24 24' width='20' xmlns='http://www.w3.org/2000/svg'><path d='M7 10l5 5 5-5z'/></svg>");
        background-position: right 0px bottom 10px;
        background-repeat: no-repeat;
        background-size: 20px;
    }

    button {
        width: 100%;
        padding: 15px;
        background: #8ebe20;
        color: white;
        border: none;
        border-radius: 12px;
        font-size: 17px;
        cursor: pointer;
        transition: 0.25s;
        margin-top: 10px;
    }

    button:hover {
        background: #8ebe20;
    }

    /* --- Mobile responsive adjustments --- */
    @media (max-width: 480px) {
        .form-container {
            padding: 25px;
        }
        button {
            padding: 14px;
        }
    }
</style>
</head>

<body>

<div class="form-container">
    <h2>Review Form</h2>

    <form action="/submit" method="POST">

        <div class="input-group">
            <label for="email">Email</label>
            <input type="email" id="email" name="email" placeholder="you@example.com" required>
        </div>

        <div class="input-group">
            <label for="name">Name</label>
            <input type="text" id="name" name="name" placeholder="Your full name" required>
        </div>

        <div class="input-group">
            <label for="location">Joined in from</label>
            <select id="location" name="location" required>
                <option value="">Select location</option>
                <option value="NG">Nigeria (NG)</option>
                <option value="KE">Kenya (KE)</option>
                <option value="GH">Ghana (GH)</option>
                <option value="ZA">South Africa (ZA)</option>
            </select>
        </div>

        <div class="input-group">
            <label for="review">What do you think of the VPS new templates</label>
            <select id="review" name="review" required>
                <option value="Excellent Addition">Excellent Addition</option>
                <option value="Phenomenal">Phenomenal</option>
                <option value="Good">Good</option>
                <option value="Meh">Meh</option>
                <option value="Eh.">Eh.</option>
                <option value="Hmm">Hmm</option>
                <option value="Wait… what are we talking about again?">Wait… what are we talking about again?</option>
            </select>
        </div>

        <button type="submit">Submit Review</button>
    </form>
</div>

</body>
</html>
`
	w.Write([]byte(response))
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	input = Attendee{
		Email:    r.FormValue("email"),
		Name:     r.FormValue("name"),
		Location: r.FormValue("location"),
		Review:   r.FormValue("review"),
	}
	
	if err = insertAttendee(conn, input); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("sql error: failed to write data for %s - %s", input.Name, err)
		return
	}

	http.Redirect(w, r, "/thx", http.StatusMovedPermanently)

}

func thxHandler(w http.ResponseWriter, r *http.Request) {
    if input.Email == "" {
		w.WriteHeader(http.StatusForbidden)
        http.Redirect(w, r, "/", http.StatusMovedPermanently)
        return
    }

	firstname := strings.Split(input.Name, " ")[0]
    response := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Thank You</title>

<style>
    body {
        font-family: "Inter", sans-serif;
        background: linear-gradient(135deg, #f0f4e6, #e8f3d7);
        margin: 0;
        padding: 0;
        display: flex;
        justify-content: center;
        align-items: center;
        min-height: 100vh;
    }

    .thank-box {
        background: white;
        padding: 40px 35px;
        width: 90%%;
        max-width: 420px;
        border-radius: 18px;
        text-align: center;
        box-shadow: 0 8px 28px rgba(0,0,0,0.12);
        animation: fadeIn 0.5s ease;
        border-top: 6px solid #8ebe20;
    }

    @keyframes fadeIn {
        from { opacity: 0; transform: translateY(10px); }
        to   { opacity: 1; transform: translateY(0); }
    }

    h2 {
        font-size: 24px;
        margin-bottom: 12px;
        color: #333;
        font-weight: 600;
    }

    p {
        color: #555;
        font-size: 16px;
        line-height: 1.6;
        margin-bottom: 30px;
    }

    .icon {
        font-size: 55px;
        color: #8ebe20;
        margin-bottom: 15px;
    }

    .btn {
        display: inline-block;
        padding: 14px 22px;
        background: #8ebe20;
        color: white;
        text-decoration: none;
        border-radius: 12px;
        font-size: 16px;
        transition: 0.25s;
    }

    .btn:hover {
        background: #7aad1d;
    }
</style>
</head>

<body>

<div class="thank-box">

    <div class="icon">✔</div>

    <h2>Thank You!</h2>

    <p>
        Thank you %s for attending and sharing your review.<br>
        We truly appreciate your time and feedback.
    </p>

    <a href="/" class="btn">Return Home</a>
</div>

</body>
</html>
`, firstname)
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(response))
}