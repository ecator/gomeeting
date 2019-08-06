
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
            start_time:app.startTime,
            end_time:app.endTime,
            room_id:app.roomID,
            memo:app.memo.toString()
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
                    if (response.data.results.length>0){
                        app.roomID=response.data.results[0].id
                    }
                }
            } else {
                alert("unknown error")
            }
        })

}

// get meetings and set it to app
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

// init vue app
let app = new Vue({
    el: "#app",
    data: {
        meetings: [],
        makeDay: getToDayAsStr(),
        makeDayDef: getToDayAsStr(),
        profile: {},
        memo:"",
        startTime:8*60,   // 8:00
        endTime:17*60+30,   //17:30
        rooms:[],
        roomID:-1
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
        postMeeting(){
            if (confirm("Add a meeting?")){
                addMeeting()
            }
        }
    },
    watch: {
        makeDay(v) {
            getMeetings()
        }
    },
    computed:{
        canAdd(){
            if (this.roomID>0 && this.startTime>0 && this.endTime>0 && this.memo!="" && this.makeDay!=""){
                return false
            }else{
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
        .then(()=>getRooms())

// init layDate
laydate.render({
    elem: '#makeDay',
    lang: "en",
    value: app.makeDayDef,
    showBottom: true,
    btns:["now"],
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
    value: "08:00 - 17:30",
    showBottom: true,
    done: function (v, d, end) {
        //console.log(v)
        let s=v.split(" - ")
        let s1=s[0].split(":")
        let s2=s[1].split(":")
        app.startTime=parseInt(s1[0])*60+parseInt(s1[1])
        app.endTime=parseInt(s2[0])*60+parseInt(s2[1])
    }
})