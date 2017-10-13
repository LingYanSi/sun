// 获取设备类型，进行css fix
; (function xx() {
    const UA = navigator.userAgent.toLowerCase()

    const version = {
        ios() {
            const v = (UA.match(/os\s*((\d_?)+)/i) || [])[1]
            return v.replace(/_/g, '.')
        },
        android() {
            const v = (UA.match(/android\s*((\d.?)+)/i) || [])[1]
            return v
        },
        ie() {
            const v = (UA.match(/MSIE\s*((\d.?)+)/i) || [])[1]
            return v
        },
        wechat() {
            const v = (UA.match(/micromessenger\/((\d.?)+)/i) || [])[1]
            return v.replace(/\.(\d+)/g, ($0, $1) => (+$1 > 10 ? $1 : `0${$1}`))
        },
    }

    const os = {
        ios: UA.match(/iphone|ipad|ipod/i) && {
            v: version.ios(),
        },
        ipad: !!UA.match(/ipad/i),
        iphone: !!UA.match(/iphone/i),
        android: UA.match(/android/i) && {
            v: version.android(),
        },
        webkit: UA.match(/webkit/i),
        ie: UA.match(/MSIE/i) && {
            v: version.ie(),
        },
        wechat: UA.match(/micromessenger/i) && {
            v: version.wechat(),
        },
    }
    os.phone = os.ios || os.android
    os.pc = !os.phone

    document.querySelector('html').className += ` ${Object.keys(os).map(key => os[key] && `device-${key}`).filter(i => i).join(' ')}`
}())

// 修复当浏览器被缩放时，字体布局混乱的bug
; (function xx() {
    function setFontSize() {
        // 如果移动端就满屏显示吧
        const IS_PC = !window.navigator.userAgent.toLowerCase().match(/phone|android|ipad|ipod/)

        let maxWidth = IS_PC
            ? Math.min(750, window.innerWidth)
            : window.innerWidth
        maxWidth = Math.max(320, maxWidth)

        const fs = (maxWidth / 750) * 100
        const style = document.querySelector('html').style
        style.fontSize = `${fs}px`

        // 新增一个div，校验其clientWidth与window.innerWidth的关系，然后对fontSize做一个二次校正
        // 我想如果页面是基于vw布局的话，应该就没有这个问题了
        const div = document.createElement('div')
        div.style.width = '7.50rem'
        document.body.appendChild(div)

        // 如果字体被缩小就增大fontSize，反则缩小FontSize,
        style.fontSize = `${(fs * maxWidth) / div.clientWidth}px`

        document.body.removeChild(div)
    }

    window.addEventListener('resize', setFontSize)

    setFontSize()
}())
