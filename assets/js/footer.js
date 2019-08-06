Vue.component("my-footer",{
    template:`
            <div class="content has-text-centered">
                <p>
                    <strong>GoMeeting</strong> by
                    <a href="https://github.com/ecator" class="button is-small">
                        <span class="icon is-small">
                            <i class="fab fa-github"></i>
                        </span>
                        <span>Martin</span>
                    </a>
                </p>
            </div>
    `
})

new Vue({
    el:"#footer"
})