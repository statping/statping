import axios from 'axios'
import API from './API'
import subSeconds from 'date-fns/subSeconds'
import getUnixTime from 'date-fns/getUnixTime'

class Graphing {
    constructor () {

    }

    subtract(seconds) {
        return getUnixTime(subSeconds(new Date(), seconds))
    }

    now() {
        return getUnixTime(new Date())
    }

    async hits(service, days, group='24h') {
        const query = await API.service_hits(service.id, this.subtract(86400 * days), this.now(), group, false)
        let total = 0;
        let high = 0;
        let low = 99999999;
        query.map((d) => {
            if (high <= d.amount) {
                high = d.amount
            }
            if (low >= d.amount && d.amount !== 0) {
                low = d.amount
            }
            total += d.amount;
        });
        const average = total / query.length;
        return {chart: query, average: average, total: total, high: high, low: low}
    }

    async pings(service, days, group='24h') {
        const query = await API.service_ping(service.id, this.subtract(86400 * days), this.now(), group)
        let total = 0;
        let high = 0;
        let low = 99999999;
        query.map((d) => {
            if (high <= d.amount) {
                high = d.amount
            }
            if (low >= d.amount && d.amount !== 0) {
                low = d.amount
            }
            total += d.amount;
        });
        const average = total / query.length;
        return {chart: query, average: average, total: total, high: high, low: low}
    }

    async failures(service, days, group='24h') {
        const query = await API.service_failures_data(service.id, this.subtract(86400 * days), this.now(), group)
        let total = 0;
        let high = 0;
        let lowest = 99999999;
        query.map((d) => {
            if (d.amount >= high) {
                high = d.amount
            }
            if (lowest >= d.amount) {
                lowest = d.amount
            }
            total += d.amount;
        });
        const average = total / query.length;
        return {data: query, average: average, total: total, high: high, low: lowest}
    }

}

const graphing = new Graphing()
export default graphing
