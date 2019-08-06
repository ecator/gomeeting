let app = new Vue({
    el: "#login",
    data: {
        username: "",
        password: ""
    },
    methods: {
        login: function () {
            //console.log(this.username,this.password,md5(this.password))
            axios.post("/login", {
                username: this.username,
                password: md5(this.password)
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

        }
    },
    computed: {
        canLogin: function () {
            if (this.username != "" && this.password != "") {
                return false
            } else {
                return true
            }
        }
    }
})