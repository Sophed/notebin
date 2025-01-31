function login() {
    auth("login");
}

function register() {
    auth("register");
}

function auth(route) {
    const username = document.getElementById("inp-username").value;
    const password = document.getElementById("inp-password").value;
    if (username == "" || password == "") {
        errEmpty();
        return;
    }
    let auth = {
      username: username,
      password: password,
    };
    fetch("/api/" + route, {
        method: "POST",
        body: JSON.stringify(auth),
        headers: {
            "Content-type": "application/json; charset=UTF-8",
        },
    }).then((response) => {
        if (response.status != 200) {
            alert(response.statusText);
            return;
        }
        response.json().then((json) => {
            if (json.error != null) {
                alert(json.error);
                return;
            }
            storeToken(json.token);
        });
    });
}

const TOKEN_KEY = "notebin-auth-token"

function storeToken(token) {
    localStorage.setItem(TOKEN_KEY, token)
    window.location.href = "/";
}

function getToken() {
    return localStorage.getItem(TOKEN_KEY)
}

function logOut() {
    localStorage.setItem(TOKEN_KEY, "")
    window.location.href = "/";
}

function errEmpty() {
    alert("One or more fields are empty")
}