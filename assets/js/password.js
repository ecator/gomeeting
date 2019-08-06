
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
                    profile.profile.name_disabled = true
                    profile.profile.email = response.data.results.email
                    profile.profile.email_show = true
                    profile.profile.email_disabled = true
                }
            } else {
                alert("unknown error")
            }
        })
}


let password = new Vue({
    el: "#password",
    data: {
        password1:"",
        password2:""
    },
    methods: {
        modify: function () {
            if (!confirm("Change your password?")) return
            axios.put("/api/user/"+profile.profile.id, {
                password: md5(this.password1) ,
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
    },
    computed: {
        canModify: function () {
            if (this.password1==this.password2 && this.password1!="" && this.password1.length>=6) {
                return false
            } else {
                return true
            }
        }
    }
})

// init data
getUser()