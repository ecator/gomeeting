//  export a global function to pull Metadata
function pullMetadata() {
    return new Promise((resolve, reject) => {
        axios.get("/meta")
            .then(function (response) {
                if (response.data.hasOwnProperty("status")) {
                    if (response.data.status != 0) {
                        reject(response.data.results)
                    } else {
                        resolve(response.data.results)
                    }
                } else {
                    reject("unknown error")
                }
            })
    })
}