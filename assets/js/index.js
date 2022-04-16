
/**
 * return timestamp of today
 * @returns {Number} timestamp 
 */
function getTodayAsTimestamp() {
    let n = new Date()
    n = new Date(n.getFullYear(), n.getMonth(), n.getDate())
    let s = parseInt(n.getTime() / 1000)
    return s
}

/**
 * return date from second timestamp
 * @param {Number} t timestamp
 * @returns {Date} date 
 */
function timestamp2date(t) {
    return new Date(t * 1000)
}

/**
 * return second timestamp from date
 * @param {Date} d the date object
 * @returns {Number} timestamp
 */
function date2timestamp(d) {
    return parseInt(d.getTime() / 1000)
}

/**
 * get formatted string of today
 * @returns {String} yyyy-mm-dd
 */
function getTodayAsStr() {
    let n = new Date()
    let s = n.getFullYear().toString() + "-" + ("0" + (n.getMonth() + 1)).substr(-2) + "-" + ("0" + n.getDate()).substr(-2)
    return s
}

/**
 * return timestamp of minutes
 * @returns {Number} timestamp 
 */
function getNowAsTimestampOfMins() {
    let n = new Date()
    n = new Date(n.getFullYear(), n.getMonth(), n.getDate(), n.getHours(), n.getMinutes())
    let s = date2timestamp(n)
    return s
}

/**
 * open teams scheduling dialog
 * @see https://docs.microsoft.com/en-us/microsoftteams/platform/concepts/build-and-test/deep-links#generate-a-deep-link-to-the-scheduling-dialog
 * @param {String} subject 
 * @param {String} content 
 * @param {String[]} attendees 
 * @param {Date} startTime 
 * @param {Date} endTime 
 */
function openTeamsSchedule(subject, content, attendees, startTime, endTime) {
    //console.log(`subject=${subject}, content=${content}, attendees=${attendees}, startTime=${startTime}, endTime=${endTime}`)
    let hasParam = false;
    let entrypoint = app.teams.entrypoint;
    if (subject) {
        entrypoint += `${hasParam ? '&' : '?'}subject=${encodeURIComponent(subject)}`;
        hasParam = true;
    }
    if (content) {
        entrypoint += `${hasParam ? '&' : '?'}content=${encodeURIComponent(content)}`;
        hasParam = true;
    }
    if (attendees && Array.isArray(attendees) && attendees.length > 0) {
        entrypoint += `${hasParam ? '&' : '?'}attendees=${encodeURIComponent(attendees.join(","))}`;
        hasParam = true;
    }
    if (startTime && startTime instanceof Date) {
        entrypoint += `${hasParam ? '&' : '?'}startTime=${encodeURIComponent(startTime.toISOString())}`;
        hasParam = true;
    }
    if (endTime && endTime instanceof Date) {
        entrypoint += `${hasParam ? '&' : '?'}endTime=${encodeURIComponent(endTime.toISOString())}`;
        hasParam = true;
    }
    window.open(entrypoint);
}

/**
 * get meetings
 *
 */
function getMeetings() {
    axios.get("/api/meeting", {
        params: {
            start_time: app.makeDay,
            end_time: app.makeDay + 24 * 3600 - 60
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

/**
 * add meetings
 *
 */
function addMeeting() {
    axios.post("/api/meeting", {
        create_time: date2timestamp(new Date()),
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

/**
 * get rooms
 *
 */
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

/**
 * get notification
 * 
 */
function getNotification() {
    axios.get("/api/notification")
        .then(function (response) {
            if (response.data.hasOwnProperty("status")) {
                if (response.data.status == 0) {
                    app.notification = response.data.results.message
                } else if (response.data.status == 9003) {
                    // no notification
                    app.notification = ""
                } else {
                    alert(response.data.results)
                }
            } else {
                alert("unknown error")
            }
        })

}

/**
 * delete meeting and refresh
 *
 * @param {Object} m the meeting object
 */
function delMeeting(m) {
    axios.delete("/api/meeting", {
        params: {
            meeting_id: m.id,
            maker: m.maker.id
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

/**
 * change sort flg and run sort meetings
 * @param {Element} sortEl - sort icon element
 * @param {string} m - sort icon name
 */
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

/**
 * run sort meetings
 *
 */
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
        makeDay: getTodayAsTimestamp(),
        makeDayStr: getTodayAsStr(),
        profile: {},
        memo: "",
        startTime: getNowAsTimestampOfMins(),
        endTime: getNowAsTimestampOfMins() + 3600 >= getTodayAsTimestamp() + 24 * 3600 ? getNowAsTimestampOfMins() : getNowAsTimestampOfMins() + 3600,
        rooms: [],
        roomID: -1,
        sortFlgs: [{ id: "maker", flg: "down" }, { id: "end_time", flg: "down" }, { id: "start_time", flg: "down" }, { id: "room", flg: "down" }],
        notification: "",
        notificationInput: "",
        teams: { enable: false, entrypoint: "https://teams.microsoft.com/l/meeting/new" }
    },
    methods: {
        openTeamsSch(m) {
            let sub = m.memo
            let content = `${m.room.name}\n${sub}`
            let startTime = timestamp2date(m.start_time)
            let endTime = timestamp2date(m.end_time)
            openTeamsSchedule(sub, content, [], startTime, endTime)
        },
        deleteMeeting(m) {
            if (confirm("Delete this meeting?")) {
                delMeeting(m)
            }
        },
        extendMeeting(m) {
            this.roomID = m.room.id
            this.memo = `forked from ${m.room.name}(${this.timestamp2hm(m.start_time)} - ${this.timestamp2hm(m.end_time)})`
            let startTime = timestamp2date(m.end_time)
            let endTime = timestamp2date(m.end_time)
            let ctl_time = document.getElementById("timeRange")
            let ctl_memo = document.getElementById("memo")
            let ms = this.meetings.filter((i) => i.room.id == m.room.id && i.start_time >= m.start_time)
            ms.sort((a, b) => a['start_time'] - b['start_time'])
            let msg = ""
            if (ms.length == 1) {
                startTime = timestamp2date(ms[0].end_time)
                let diffsec = this.makeDay + 24 * 3600 - 60 - date2timestamp(startTime)
                if (diffsec >= 30 * 60) {
                    diffsec = 30 * 60
                    msg = "You can extend this meeting by at least 30 minutes."
                } else {
                    msg = `You can extend this meeting by up to ${parseInt(diffsec / 60)} minutes.`
                }
                endTime = timestamp2date(date2timestamp(startTime) + diffsec)
            } else if (ms[1].start_time - ms[0].end_time >= 30 * 60) {
                startTime = timestamp2date(ms[0].end_time)
                endTime = timestamp2date(date2timestamp(startTime) + 30 * 60)
                msg = "You can extend this meeting by at least 30 minutes."
            } else if (ms[1].start_time - ms[0].end_time >= 10 * 60) {
                startTime = timestamp2date(ms[0].end_time)
                endTime = timestamp2date(ms[1].start_time)
                msg = `You can extend this meeting by up to ${parseInt((date2timestamp(endTime) - date2timestamp(startTime)) / 60)} minutes.`
            } else {
                msg = "Sorry,you can't extend this meeting because of another one coming up!\nYou may try to change a room."
            }
            this.startTime = date2timestamp(startTime)
            this.endTime = date2timestamp(endTime)
            ctl_time.value = this.timestamp2hm(this.startTime) + " - " + this.timestamp2hm(this.endTime)
            ctl_memo.focus()
            setTimeout(() => ctl_memo.select(), 100)
            alert(msg)
        },
        /**
         * convert timestamp to hh:mm
         *
         * @param {Number} t timestamp
         * @returns {string} formatted string of hh:mm 
         */
        timestamp2hm(t) {
            let d = timestamp2date(t)
            return `0${d.getHours()}`.substr(-2) + ":" + `0${d.getMinutes()}`.substr(-2)
        },
        /**
         * add new meeting
         *
         */
        postMeeting() {
            if (confirm("Add a meeting?")) {
                addMeeting()
            }
        }
        ,
        /**
         * check if can delete one meeting
         * @param {Object} m meeting object
         * @returns {boolean} 
         */
        canDel(m) {
            if (this.profile.id == 0) {
                return true
            } else if (this.profile.id == m.maker.id || parseInt(this.profile.level) < parseInt(m.maker.level)) {
                let startTime = m.start_time
                if (date2timestamp(new Date()) - 5 * 60 > startTime) {
                    return false
                }
                return true
            } else {
                return false
            }
        },
        /**
         * check if cross days with makeday
         * @param {number} t
         * @returns {boolean}
         */
        isCrossDay(t) {
            let makeDay = timestamp2date(this.makeDay)
            let target = timestamp2date(t)
            return makeDay.getDate() != target.getDate()
        },
        delNotification() {
            if (confirm("Delete this notification?")) {
                axios.delete("/api/notification")
                    .then(() => getNotification())
                    .then(setTimeout(() => this.notificationRowAdjust(), 500))
            }
        },
        addNotification() {
            axios.post("/api/notification", { message: this.notificationInput })
                .then(function (response) {
                    if (response.data.hasOwnProperty("status")) {
                        if (response.data.status != 0) {
                            alert(response.data.results)
                        } else {
                            getNotification()
                        }
                    } else {
                        alert("unknown error")
                    }
                })
        },
        /**
         * start input notification
         *
         */
        notificationRowAdjust() {
            let textArea = document.getElementById("notificationInput")
            let mt = textArea.value.match(/\n/g)
            if (mt) {
                textArea.rows = mt.length
            } else {
                textArea.rows = 1
            }
            // hide scroll
            while (textArea.scrollHeight > textArea.clientHeight) {
                textArea.rows++
            }
        }
    },
    watch: {
        makeDay(v) {
            getMeetings()
            let startTime = timestamp2date(this.startTime)
            let endTime = timestamp2date(this.endTime)
            let makeDay = timestamp2date(this.makeDay)
            makeDay.setHours(startTime.getHours())
            makeDay.setMinutes(startTime.getMinutes())
            app.startTime = date2timestamp(makeDay)
            makeDay.setHours(endTime.getHours())
            makeDay.setMinutes(endTime.getMinutes())
            app.endTime = date2timestamp(makeDay)
        },
        notification(v) {
            if (v != "") {
                this.notificationInput = v
            }
        },
        notificationInput(v) {
            this.notificationRowAdjust()
        }
    },
    computed: {
        canAdd() {
            if (this.roomID > 0 && this.startTime > 0 && this.endTime > 0 && this.memo != "" && this.makeDay != "") {
                return false
            } else {
                return true
            }
        },
        parseNotification() {
            return markdown.toHTML(this.notification)
        },
        canAddNotification() {
            return this.notificationInput ? true : false
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
                app.profile.ldap = response.data.results.ldap
            }
        } else {
            alert("unknown error")
        }
    }).then(() => getMeetings())
    .then(() => getRooms())
    .then(() => getNotification())

// init layDate
laydate.render({
    elem: '#makeDay',
    lang: "en",
    value: app.makeDayStr,
    showBottom: true,
    btns: ["now"],
    done: function (v, d, end) {
        //console.log(v,d)
        app.makeDay = date2timestamp(new Date(d.year, d.month - 1, d.date))
    }
})

laydate.render({
    elem: '#timeRange',
    type: "time",
    range: true,
    format: "HH:mm",
    lang: "en",
    value: `${app.timestamp2hm(app.startTime)} - ${app.timestamp2hm(app.endTime)}`,
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
        let makeDay = timestamp2date(app.makeDay)
        makeDay.setHours(s1[0])
        makeDay.setMinutes(s1[1])
        app.startTime = date2timestamp(makeDay)
        makeDay.setHours(s2[0])
        makeDay.setMinutes(s2[1])
        app.endTime = date2timestamp(makeDay)
    }
})

pullMetadata().then(meta => {
    if (meta.teams.enable) {
        app.teams.enable = meta.teams.enable
        app.teams.entrypoint = meta.teams.entrypoint || app.teams.entrypoint
    }
}, err => {
    console.log("pullMetadata error: " + err)
})