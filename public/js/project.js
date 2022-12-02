let dataBlog = [];

function addBlog(event){
    event.preventDefault()

    let inputblogName = document.getElementById("blog-name").value;
    let inputStartDate = document.getElementById("date-start").value;
    let inputEndDate = document.getElementById("date-end").value;
    let inputCheckbox = document.getElementsByName("inputCheck").value;
    let inputBlogDesc = document.getElementById("blog-description").value;
    let inputUploadImage = document.getElementById("upload-image").files[0];

    inputUploadImage = URL.createObjectURL(inputUploadImage);

    let blog = {

            inputblogName,
            inputStartDate,
            inputEndDate,
            inputBlogDesc,
            inputCheckbox,
            inputUploadImage,
            postAt: new Date(),
            author: "Try Widodo Putra"
        }

        dataBlog.push(blog);
        console.log(dataBlog)
        renderBlog()
    }

// displaying blog
function renderBlog() {

    document.getElementById("contents").innerHTML = ''

    for (let index = 0; index < dataBlog.length; index++) {
        console.log("test",dataBlog[index])

        document.getElementById("contents").innerHTML +=`
        <article class="blog-item">
                    <img src="${dataBlog[index].inputUploadImage}" alt="">
                    <div class="blog-name">
                    <h3>
                        <a href="blog-detail.html" target="_blank">
                            ${dataBlog[index].inputblogName}
                        </a>
                    </h3>
                        <div class="blog-duration">
                            <p>Duration: ${getDistanceTime(dataBlog[index].postAt)}</p>
                            <p><b>Created</b> ${dataBlog[index].inputStartDate}</p>
                            <p><b>End Date:</b> ${dataBlog[index].inputEndDate}</p>
                        </div>
                    </div>
                    <div class="blog-description">
                        <p>
                          ${dataBlog[index].inputBlogDesc}
                        </p>
                        <div class="blog-tech-info">
                            <img src="./assets/icons/nodejs.svg">
                            <img src="./assets/icons/react-native.svg">
                            <img src="./assets/icons/nextjs.svg">
                            <img src="./assets/icons/typescript.svg">
                        </div>
                    </div>
                    <div class="action-btn">
                        <button class="btn btn-primary">edit</button>
                        <button class="btn btn-white">delete</button>
                    </div>
                </article>
        `
        
    }
}

function getFullTime(time) {

    let monthName = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec']

    let date = time.getDate()
    console.log(date)

    let monthIndex = time.getMonth()
    console.log(monthIndex)

    let year = time.getFullYear()
    console.log(year)

    let hours = time.getHours()
    let minutes = time.getMinutes()

    console.log(hours)

    if (hours <= 9) {
        hours = "0" + hours
    } 
    
    if (minutes <= 9) {
        minutes = "0" + minutes
    }

    // 14 Oct 2022 09:07 WIB
    return `${date} ${monthName[monthIndex]} ${year} ${hours}:${minutes} WITA`
}

function getDistanceTime(time) {
    let timeNow = new Date()
    let timePost = time

    let distance = timeNow - timePost //milisecond
    console.log(distance)

    let milisecond = 1000 
    let secondInHours = 3600 
    let hoursInDay = 24 

    let distanceDay = Math.floor(distance / (milisecond * secondInHours * hoursInDay))
    let distanceHours = Math.floor(distance / (milisecond * 60 * 60))
    let distanceMinutes = Math.floor(distance / (milisecond * 60))
    let distanceSecond = Math.floor(distance / milisecond)

    if (distanceDay > 0) {
        return `${distanceDay} hari yang lalu`
    } else if (distanceHours > 0) {
        return `${distanceHours} jam yang lalu`
    } else if (distanceMinutes > 0) {
        return `${distanceMinutes} menit yang lalu`
    } else {
        return `${distanceSecond} detik yang lalu`
    }
}


setInterval(function() {
    renderBlog()
}, 60000)

function getDurationTime (startDate, endDate) {
    let date1 = new Date(startDate);
    let date2 = new Date(endDate);
    let diffDate = Math.abs(date2 - date1);
    let projectDuration = Math.floor(diffDate / (1000 * 3600 * 24));

    let calculateDuration = "";
    let durationTotal = "";

    if (projectDuration > 30) {
        // not calculate correctly
        calculateDuration = Math.round(projectDuration / 30);
        durationTotal = `${calculateDuration} month(s)`;
    } else {
        durationTotal = `${projectDuration} day(s)`;
    }
    return durationTotal;
}

