Vue.component("my-nav", {
    props: ["isAdmin","isLdap"],
    template: `
    <nav class="navbar is-primary" role="navigation" aria-label="main navigation">
            <div class="navbar-brand">
                <a class="navbar-item" href="/">
                    GoMeeting
                </a>
            </div>

            <div id="navbarMenu" class="navbar-menu is-active">
                <div class="navbar-start">
                    <a class="navbar-item" href="/user">
                        Profile
                    </a>
                    <a v-if="!isLdap" class="navbar-item" href="/password">
                        Password
                    </a>
                    <a v-if="isAdmin" class="navbar-item" href="/admin">
                        Admin
                    </a>
                </div>

                <div class="navbar-end">
                    <div class="navbar-item">
                            <a class="button is-light" href="/logout">
                                Logout
                            </a>
                    </div>
                </div>
            </div>
        </nav>
    
    `
})

axios.get("/api/user/my")
    .then(function (response) {
        if (response.data.hasOwnProperty("status")) {
            if (response.data.status != 0) {
                alert(response.data.results)
            } else {
                new Vue({
                    el: "#nav",
                    data: {
                        isAdmin: response.data.results.id == 0 ? true : false,
                        isLdap: response.data.results.ldap
                    }
                })
            }
        } else {
            alert("unknown error")
        }
    })