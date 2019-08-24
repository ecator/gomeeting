
// return yyyymmdd
function getToDayAsInt() {
    let n = new Date
    let s = n.getFullYear().toString() + ("0" + (n.getMonth() + 1)).substr(-2) + ("0" + n.getDate()).substr(-2)
    return parseInt(s)
}

// return yyyy-mm-dd
function getToDayAsStr() {
    let n = new Date
    let s = n.getFullYear().toString() + "-" + ("0" + (n.getMonth() + 1)).substr(-2) + "-" + ("0" + n.getDate()).substr(-2)
    return s
}

// return hours*60 + mins
function getNowAsMins() {
    let n = new Date()
    let s = n.getHours() * 60 + n.getMinutes()
    return s
}

// get meetings
function getMeetings() {
    axios.get("/api/meeting", {
        params: {
            make_date: app.makeDay ? app.makeDay.replace(/-/g, "") : app.makeDayDef.replace(/-/g, "")
        }
    })
        .then(function (response) {
            if (response.data.hasOwnProperty("status")) {
                if (response.data.status != 0) {
                    alert(response.data.results)
                } else {
                    app.meetings = app.meetings.filter(() => false)
                    for (m of response.data.results) {
                        app.meetings.push(m)
                    }
                    runSortMeetings()
                }
            } else {
                alert("unknown error")
            }
        })

}

// add meeting
function addMeeting() {
    axios.post("/api/meeting", {
        make_date: parseInt(app.makeDay ? app.makeDay.replace(/-/g, "") : app.makeDayDef.replace(/-/g, "")),
        start_time: app.startTime,
        end_time: app.endTime,
        room_id: app.roomID,
        memo: app.memo.toString()
    })
        .then(function (response) {
            if (response.data.hasOwnProperty("status")) {
                if (response.data.status != 0) {
                    alert(response.data.results)
                } else {
                    // add ok
                    getMeetings()
                }
            } else {
                alert("unknown error")
            }
        })

}

// get rooms
function getRooms() {
    axios.get("/api/rooms")
        .then(function (response) {
            if (response.data.hasOwnProperty("status")) {
                if (response.data.status != 0) {
                    alert(response.data.results)
                } else {
                    app.rooms = app.rooms.filter(() => false)
                    for (r of response.data.results) {
                        app.rooms.push(r)
                    }
                    if (response.data.results.length > 0) {
                        app.roomID = response.data.results[0].id
                    }
                }
            } else {
                alert("unknown error")
            }
        })

}

// delete meeting and refresh
function delMeeting(m) {
    axios.delete("/api/meeting", {
        params: {
            make_date: m.make_date,
            room_id: m.room.id,
            start_time: m.start_time,
            end_time: m.end_time,
            maker: app.profile.id
        }
    })
        .then(function (response) {
            if (response.data.hasOwnProperty("status")) {
                if (response.data.status != 0) {
                    alert(response.data.results)
                } else {
                    getMeetings()
                }
            } else {
                alert("unknown error")
            }
        })

}

// change sort flg and run sort meetings
function sortMeetings(sortEl, m) {
    // find flg and shift position
    let flg = ""
    for (let i = 0; i < app.sortFlgs.length; i++) {
        if (app.sortFlgs[i].id == m) {
            if (app.sortFlgs[i].flg == "down") {
                app.sortFlgs[i].flg = "up"
            } else {
                app.sortFlgs[i].flg = "down"
            }
            flg = app.sortFlgs[i].flg
            // shift position
            let tmp = app.sortFlgs[i]
            app.sortFlgs.splice(i, 1)
            app.sortFlgs.push(tmp)
            break
        }
    }
    sortEl.setAttribute("class", "fas fa-caret-" + flg)
    runSortMeetings()
}
// run sort meetings
function runSortMeetings() {
    for (let i of app.sortFlgs) {
        //console.log("sort",i.id,i.flg)
        switch (i.id) {
            case "maker":
                if (i.flg == "up") {
                    app.meetings.sort((a, b) => a.maker.username <= b.maker.username ? -1 : 1)
                } else {
                    app.meetings.sort((a, b) => a.maker.username <= b.maker.username ? 1 : -1)
                }
                break
            case "end_time":
            case "start_time":
                if (i.flg == "up") {
                    app.meetings.sort((a, b) => a[i.id] - b[i.id])
                } else {
                    app.meetings.sort((a, b) => b[i.id] - a[i.id])
                }
                break
            case "room":
                if (i.flg == "up") {
                    app.meetings.sort((a, b) => a.room.id - b.room.id)
                } else {
                    app.meetings.sort((a, b) => b.room.id - a.room.id)
                }
        }
    }
}

// init vue app
let app = new Vue({
    el: "#app",
    data: {
        meetings: [],
        makeDay: getToDayAsStr(),
        makeDayDef: getToDayAsStr(),
        profile: {},
        memo: "",
        startTime: getNowAsMins(),
        endTime: getNowAsMins() >= 23 * 60 ? getNowAsMins() : getNowAsMins() + 60,
        rooms: [],
        roomID: -1,
        sortFlgs: [{ id: "maker", flg: "down" }, { id: "end_time", flg: "down" }, { id: "start_time", flg: "down" }, { id: "room", flg: "down" }]
    },
    methods: {
        deleteMeeting(m) {
            if (confirm("Delete this meeting?")) {
                delMeeting(m)
            }
        },
        // convert mins to hh:mm
        mins2hm(mins) {
            return ("0" + parseInt(mins / 60)).substr(-2) + ":" + ("0" + (mins % 60)).substr(-2)
        },
        // add new meeting
        postMeeting() {
            if (confirm("Add a meeting?")) {
                addMeeting()
            }
        }
    },
    watch: {
        makeDay(v) {
            getMeetings()
        }
    },
    computed: {
        canAdd() {
            if (this.roomID > 0 && this.startTime > 0 && this.endTime > 0 && this.memo != "" && this.makeDay != "") {
                return false
            } else {
                return true
            }
        }
    }
})



// init data
// get user info and meetings
axios.get("/api/user/my")
    .then(function (response) {
        if (response.data.hasOwnProperty("status")) {
            if (response.data.status != 0) {
                alert(response.data.results)
            } else {
                app.profile.id = response.data.results.id
                app.profile.username = response.data.results.username
                app.profile.level = response.data.results.level
                app.profile.org = response.data.results.org.name
                app.profile.name = response.data.results.name
                app.profile.email = response.data.results.email
            }
        } else {
            alert("unknown error")
        }
    }).then(() => getMeetings())
    .then(() => getRooms())

// init layDate
laydate.render({
    elem: '#makeDay',
    lang: "en",
    value: app.makeDayDef,
    showBottom: true,
    btns: ["now"],
    done: function (v, d, end) {
        //console.log(v)
        app.makeDay = v
    }
})

laydate.render({
    elem: '#timeRange',
    type: "time",
    range: true,
    format: "HH:mm",
    lang: "en",
    value: `${app.mins2hm(app.startTime)} - ${app.mins2hm(app.endTime)}`,
    showBottom: true,
    btns: ["confirm"],
    ready: function (d) {
        //console.log(d)
        // hide seconds area
        for (let i of document.querySelectorAll(".laydate-time-list >li:nth-child(3)")) {
            i.style.display = "none"
        }
        // set width as half
        for (let i of document.querySelectorAll(".laydate-time-list >li")) {
            i.style.width = "50%"
        }
    },
    done: function (v, d, end) {
        //console.log(v)
        let s = v.split(" - ")
        let s1 = s[0].split(":")
        let s2 = s[1].split(":")
        app.startTime = parseInt(s1[0]) * 60 + parseInt(s1[1])
        app.endTime = parseInt(s2[0]) * 60 + parseInt(s2[1])
    }
})