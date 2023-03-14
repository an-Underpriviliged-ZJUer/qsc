// 定义获取弹幕数据的函数
function getDanmuData() {
    return axios.get('http://localhost:8080/danmu')
        .then(response => {
            return response.data;
        })
        .catch(error => {
            console.log(error);
        });
}

// 定义将弹幕数据添加到页面上的函数
function appendDanmu(danmu) {
    var color = danmu.rgb;
    var content = danmu.content;
    var span = document.createElement('span');
    span.style.color = color;
    span.innerText = content;
    var container = document.getElementById('danmu-container');
    container.appendChild(span);
    setTimeout(function() {
        container.removeChild(span);
    }, 10000); // 10秒后移除该弹幕
}

// 定义启动弹幕滚动的函数
function startDanmuRolling() {
    setInterval(function() {
        getDanmuData().then(danmu => {
            appendDanmu(danmu);
        });
    }, 1000); // 每1秒获取一次弹幕数据并添加到页面上
}

// 页面加载完成后启动弹幕滚动
window.onload = function() {
    startDanmuRolling();
};
