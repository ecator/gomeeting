

// get user info and set it to profile  and document.title
function getUser() {
    axios.get("/api/user/my")
        .then(function (response) {
            if (response.data.hasOwnProperty("status")) {
                if (response.data.status != 0) {
                    alert(response.data.results)
                } else {
                    document.title = "GoMeeting-" + response.data.results.name
                    profile.profile.id = response.data.results.id
                    profile.profile.id_show = true
                    profile.profile.username = response.data.results.username
                    profile.profile.username_show = true
                    profile.profile.org.name = response.data.results.org.name
                    profile.profile.org.name_show = true
                    profile.profile.name = response.data.results.name
                    profile.profile.name_show = true
                    profile.profile.name_disabled = false
                    profile.profile.email = response.data.results.email
                    profile.profile.email_show = true
                    profile.profile.email_disabled = false
                }
            } else {
                alert("unknown error")
            }
        })
        .then(()=>canModify())
}

function canModify() {
    let btn_modify=document.getElementById("btn_modify")
    if (profile.profile.name != "" && (new RegExp(/\w+@\w+\.\w+/)).test(profile.profile.email)) {
        btn_modify.disabled=false
    } else {
        btn_modify.disabled=true
    }
}

// modify user info
function modify() {
    if (!confirm("Modify your profile?")) return
    axios.put("/api/user/" + profile.profile.id, {
        name: profile.profile.name,
        email: profile.profile.email
    })
        .then(function (response) {
            //console.log(response.data)
            if (response.data.hasOwnProperty("status")) {
                if (response.data.status != 0) {
                    // ng
                    alert(response.data.results)
                } else {
                    //ok
                    alert("success")
                }
            } else {
                // unknown error
                alert("fail")
            }
        })
        .catch(function (error) {
            console.error(error)
        })
}


// init data
getUser()
