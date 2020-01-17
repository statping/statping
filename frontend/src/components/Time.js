class Time {
    now () {
        return new Date();
    }
    utc () {
        return new Date().getUTCDate();
    }
    utcToLocal (utc) {
        let u = new Date().setUTCDate(utc)
        return u.toLocaleString()
    }
}
const time = new Time()
export default time
