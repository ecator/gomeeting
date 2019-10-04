// get objs
function getObjs(t = "users") {
    let obj, pusho
    switch (t) {
        case "users":
            obj = user
            break
        case "orgs":
            obj = org
            break
        case "rooms":
            obj = room
            break
        default:
            alert("getObjs : target error")
            return
    }
    return axios.get(`/api/${t}`)
        .then(function (response) {
            if (response.data.hasOwnProperty("status")) {
                if (response.data.status != 0) {
                    alert(response.data.results)
                } else {
                    obj.list = []
                    for (let r of response.data.results) {
                        if (r.id == 0) {
                            continue
                        }

                        pusho = getInitObj(t.substring(0, t.length - 1))
                        obj.list.push(Object.assign(pusho, r))
                    }
                    // NEW OBJ
                    pusho = getInitObj(t.substring(0, t.length - 1))
                    obj.list.push(pusho)
                    obj.prepareModify(obj.list[0])
                    if (t=="orgs" && obj.list.length>=2){
                        user.orgs=obj.list.slice(0,-1)
                    }
                }
            } else {
                alert("unknown error")
            }
        })
}

// modify obj
function modObj(t) {
    let obj, data
    switch (t) {
        case "user":
            obj = user
            data = {
                level: obj.activeObj.level,
                password: obj.activeObj.password,
                name: obj.activeObj.name,
                org_id: obj.activeObj.org.id,
                email: obj.activeObj.email
            }
            break
        case "org":
            obj = org
            data = { name: obj.activeObj.name }
            break
        case "room":
            obj = room
            data = { name: obj.activeObj.name }
            break
        default:
            alert("modObj : target error")
            return
    }

    // input check
    if (obj.activeObj.hasOwnProperty('level') && obj.activeObj.level == "") {
        alert("Level must be a number")
        return
    }
    if (obj.activeObj.hasOwnProperty('name') && obj.activeObj.name == "") {
        alert("Name can not be empty")
        return
    }
    if (obj.activeObj.hasOwnProperty('email') && obj.activeObj.email == "") {
        alert("Email can not be empty")
        return
    }
    if (obj.activeObj.hasOwnProperty('password') && obj.activeObj.password != "") {
        if (!confirm("Modify password ?")) return
    }

    // confirm
    if (!confirm("Modify " + obj.activeObj.name + " ?")) return

    axios.put(`/api/${t}/${obj.activeObj.id}`, data)
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
// add a new obj
function addObj(t) {
    let obj, data
    switch (t) {
        case "user":
            obj = user
            data = {
                username:obj.activeObj.username,
                level: obj.activeObj.level,
                password: obj.activeObj.password,
                name: obj.activeObj.name,
                org_id: obj.activeObj.org.id,
                email: obj.activeObj.email
            }
            break
        case "org":
            obj = org
            data = { name: obj.activeObj.name }
            break
        case "room":
            obj = room
            data = { name: obj.activeObj.name }
            break
        default:
            alert("addObj : target error")
            return
    }
    if (obj.activeObj.hasOwnProperty('username') && obj.activeObj.username == "") {
        alert("Username must be a number")
        return
    }
    if (obj.activeObj.hasOwnProperty('password') && obj.activeObj.password == "") {
        alert("Password must be a number")
        return
    }
    if (obj.activeObj.hasOwnProperty('level') && obj.activeObj.level == "") {
        alert("Level must be a number")
        return
    }
    if (obj.activeObj.hasOwnProperty('name') && obj.activeObj.name == "") {
        alert("Name can not be empty")
        return
    }
    if (obj.activeObj.hasOwnProperty('email') && obj.activeObj.email == "") {
        alert("Email can not be empty")
        return
    }

    if (!confirm("Add " + obj.activeObj.name + " ?")) return

    axios.post(`/api/${t}`, data)
        .then(function (response) {
            //console.log(response.data)
            if (response.data.hasOwnProperty("status")) {
                if (response.data.status != 0) {
                    // ng
                    alert(response.data.results)
                } else {
                    //ok
                    getObjs(t+'s').then(function () {
                        obj.prepareModify(obj.list[obj.list.length - 2])
                        alert("success")
                    })
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

// delete a user/org/room
function delObj(t) {
    let obj
    switch (t) {
        case "user":
            obj = user
            break
        case "org":
            obj = org
            break
        case "room":
            obj = room
            break
        default:
            alert("delObj : target error")
            return
    }
    if (!confirm("Delete " + obj.activeObj.name + " ?")) return
    axios.delete("/api/" + t + "/" + obj.activeObj.id)
        .then(function (response) {
            //console.log(response.data)
            if (response.data.hasOwnProperty("status")) {
                if (response.data.status != 0) {
                    // ng
                    alert(response.data.results)
                } else {
                    //ok
                    getObjs(t + 's').then(() => {
                        obj.prepareModify(obj.list[0])
                        alert("success")
                    }
                    )
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

let tabs = new Vue({
    el: "#tabs",
    data: {
        userActive: true,
        orgActive: false,
        roomActive: false
    },
    methods: {
        activateTab(tabName) {
            this.userActive = false
            user.isActive = false
            this.orgActive = false
            org.isActive = false
            this.roomActive = false
            room.isActive = false
            switch (tabName) {
                case "user":
                    this.userActive = true
                    user.isActive = true
                    break;
                case "org":
                    this.orgActive = true
                    org.isActive = true
                    break;
                case "room":
                    this.roomActive = true
                    room.isActive = true
                    break;
                default:
                    this.userActive = true
                    user.isActive = true
                    break;
            }
        }
    }
})

// create vue obj for user/org/room
function getVueObj(el, data, methods, watch) {
    return {
        el,
        data: Object.assign({
            list: [],
            activeObj: { id: -1, name: "" },
            filter: "",
            isActive: false
        }, data),
        methods: Object.assign({
            prepareModify(o) {
                this.activeObj = o
                for (let p of this.list) {
                    p.isActive = false
                    if (p.id == o.id) {
                        p.isActive = true
                    }
                }
            }
        }, methods),
        watch: Object.assign({
            filter(n, o) {
                for (p of this.list) {
                    if ((p.id + p.name).indexOf(n) > -1 || p.id == -1) {
                        p.isShow = true
                    } else {
                        p.isShow = false
                    }
                }
            }
        }, watch)
    }
}
// return init obj struct
function getInitObj(t) {
    switch (t) {
        case "user":
            return {
                id: -1,
                username: "",
                password: "",
                level: 10,
                org: {
                    id: 1000,
                    name: ""
                },
                name: "",
                email: "",
                isActive: false,
                isShow: true
            }
        case "org":
        case "room":
            return {
                id: -1,
                name: "",
                isActive: false,
                isShow: true
            }
        default:
            alert("getInitObj : target error")
            return
    }
}

let user = new Vue(getVueObj('#user', {
    activeObj: getInitObj("user"),
    orgs: [],
    isActive: tabs.userActive
}))

let org = new Vue(getVueObj('#org', {
    activeObj: getInitObj("org"),
    isActive: tabs.orgActive
}))

let room = new Vue(getVueObj('#room', {
    activeObj: getInitObj("room"),
    isActive: tabs.roomActive
}))

getObjs('orgs').then(() => getObjs('users'))
getObjs('rooms')