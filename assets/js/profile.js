Vue.component("user-profile",{
    props:['profile'],
    template:`
    <div class="container">
        <div v-if="profile.id_show" class="field">
            <label class="label">ID</label>
            <div class="control">
                <input class="input" type="text" v-model="profile.id" placeholder="id" v-bind:disabled="profile.id_disabled" >
            </div>
        </div>
        <div v-if="profile.username_show" class="field">
            <label class="label">UserName</label>
            <div class="control">
                <input class="input" type="text" v-model="profile.username" placeholder="username" v-bind:disabled="profile.username_disabled" >
            </div>
        </div>
        <div v-if="profile.level_show" class="field">
            <label class="label">Level</label>
            <div class="control">
                <input class="input" type="text" v-model.number="profile.level" placeholder="level" v-bind:disabled="profile.level_disabled" >
            </div>
        </div>
        <div v-if="profile.org.id_show" class="field">
            <label class="label">OrgID</label>
            <div class="control">
                <input class="input" type="text" v-model.number="profile.org.id" placeholder="org id" v-bind:disabled="profile.org.id_disabled" >
            </div>
        </div>
        <div v-if="profile.org.name_show" class="field">
            <label class="label">OrgName</label>
            <div class="control">
                <input class="input" type="text" v-model="profile.org.name" placeholder="org name" v-bind:disabled="profile.org.name_disabled" >
            </div>
        </div>
        <div v-if="profile.name_show" class="field">
            <label class="label">Name</label>
            <div class="control">
                <input class="input" type="text" v-model="profile.name" placeholder="name" v-bind:disabled="profile.name_disabled" >
            </div>
        </div>
        <div v-if="profile.email_show" class="field">
            <label class="label">Email</label>
            <div class="control">
                <input class="input" type="text" v-model="profile.email" placeholder="email" v-bind:disabled="profile.email_disabled" >
            </div>
        </div>
    </div>
    `
})


let profile=new Vue({
    el:"#profile",
    data:{
        profile:{
            id:0,
            id_show:false,
            id_disabled:true,
            username:"",
            username_show:false,
            username_disabled:true,
            level:0,
            level_show:false,
            level_disabled:true,
            org:{
                id:0,
                id_show:false,
                id_disabled:true,
                name:"",
                name_show:false,
                name_disabled:true
            },
            name:"",
            name_show:false,
            name_disabled:true,
            email:"",
            email_show:false,
            email_disabled:true
        }
    }
})
