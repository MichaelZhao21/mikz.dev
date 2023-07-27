// Sets the cookie variable
function setCookie(key, val) {
    const d = new Date();
    d.setTime(d.getTime() + (1000 * 24 * 60 * 60 * 1000));
    let expires = "expires=" + d.toUTCString();
    document.cookie = key + "=" + val + ";" + expires + ";path=/";
}

// Deletes a cookie
function deleteCookie(key) {
    document.cookie = key + "=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
}
