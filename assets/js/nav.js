Vue.component("my-nav",{
    template:`
    <nav class="navbar is-primary" role="navigation" aria-label="main navigation">
            <div class="navbar-brand">
                <a class="navbar-item" href="/">
                    GoMeeting
                </a>
            </div>

            <div id="navbarBasicExample" class="navbar-menu">
                <div class="navbar-start">
                    <a class="navbar-item" href="/user">
                        Profile
                    </a>
                    <a class="navbar-item" href="password">
                        Password
                    </a>
                </div>

                <div class="navbar-end">
                    <div class="navbar-item">
                            <a class="button is-light" href="/logout">
                                Log out
                            </a>
                    </div>
                </div>
            </div>
        </nav>
    
    `
})

new Vue({
    el:"#nav"
})