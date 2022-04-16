let app = new Vue({
    el: "#login",
    data: {
        username: "",
        password: "",
        placeholder: { username: "username", password: "password" }
    },
    methods: {
        login: function () {
            //console.log(this.username,this.password,md5(this.password))
            axios.post("/login", {
                username: this.username,
                password: this.password  // to support ldap password changes to plain text
            })
                .then(function (response) {
                    //console.log(response.data)
                    if (response.data.hasOwnProperty("status")) {
                        if (response.data.status != 0) {
                            //ng
                            alert(response.data.results)
                        } else {
                            //ok
                            location.href = "/"
                        }
                    } else {
                        // unknown error
                        alert("login fail")
                    }
                })
                .catch(function (error) {
                    console.error(error)
                })

        },
        keypress: function (e) {
            if (e.keyCode == 13 && this.canLogin) {
                this.login()
            }
        }
    },
    computed: {
        canLogin: function () {
            if (this.username != "" && this.password != "") {
                return true
            } else {
                return false
            }
        }
    }
})

pullMetadata().then(meta => {
    if (meta.ldap.enable) {
        app.placeholder.username = meta.ldap.placeholder.username
        app.placeholder.password = meta.ldap.placeholder.password
    }
}, err => {
    console.log("pullMetadata error: " + err)
})