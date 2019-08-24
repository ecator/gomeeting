
// must use firefox or chrome
try {
    if (!navigator.userAgent.toLowerCase().match(/firefox|chrome/)) {
        location.href = "/brower_err"
    }
} catch (error) {
    location.href = "/brower_err"
}