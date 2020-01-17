<template>
    <form @submit="saveSettings" method="POST" action="settings">
        <div class="form-group">
            <label>Project Name</label>
            <input v-model="core.name" type="text" class="form-control" placeholder="Great Uptime">
        </div>

        <div class="form-group">
            <label>Project Description</label>
            <input v-model="core.description" type="text" class="form-control" placeholder="Great Uptime">
        </div>

        <div class="form-group row">
            <div class="col-8 col-sm-9">
                <label>Domain</label>
                <input v-model="core.domain" type="url" class="form-control">
            </div>
            <div class="col-4 col-sm-3 mt-sm-1 mt-0">
                <label class="d-inline d-sm-none">Enable CDN</label>
                <label class="d-none d-sm-block">Enable CDN</label>
                <span class="switch">
                                        <input v-model="core.using_cdn" type="checkbox" class="switch" v-bind:disabled="core.using_cdn">
                                        <label class="mt-2 mt-sm-0"></label>
                                    </span>
                <input v-model="core.using_cdn" type="hidden" name="enable_cdn" id="switch-normal-value" value="false">
            </div>
        </div>



        <div class="form-group">
            <label>Custom Footer</label>
            <textarea rows="4" class="form-control">{{core.footer}}</textarea>
            <small class="form-text text-muted">HTML is allowed inside the footer</small>
        </div>

        <div class="form-group">
            <label for="timezone">Timezone</label><span class="mt-1 small float-right">Current: {{now}}</span>
            <select class="form-control" name="timezone" id="timezone">
                <option value="-12.0" >(GMT -12:00) Eniwetok, Kwajalein</option>
                <option value="-11.0" >(GMT -11:00) Midway Island, Samoa</option>
                <option value="-10.0" >(GMT -10:00) Hawaii</option>
                <option value="-9.0" >(GMT -9:00) Alaska</option>
                <option value="-8.0" selected>(GMT -8:00) Pacific Time (US &amp; Canada)</option>
                <option value="-7.0" >(GMT -7:00) Mountain Time (US &amp; Canada)</option>
                <option value="-6.0" >(GMT -6:00) Central Time (US &amp; Canada), Mexico City</option>
                <option value="-5.0" >(GMT -5:00) Eastern Time (US &amp; Canada), Bogota, Lima</option>
                <option value="-4.0" >(GMT -4:00) Atlantic Time (Canada), Caracas, La Paz</option>
                <option value="-3.5" >(GMT -3:30) Newfoundland</option>
                <option value="-3.0" >(GMT -3:00) Brazil, Buenos Aires, Georgetown</option>
                <option value="-2.0" >(GMT -2:00) Mid-Atlantic</option>
                <option value="-1.0" >(GMT -1:00 hour) Azores, Cape Verde Islands</option>
                <option value="0.0" >(GMT) Western Europe Time, London, Lisbon, Casablanca</option>
                <option value="1.0" >(GMT +1:00 hour) Brussels, Copenhagen, Madrid, Paris</option>
                <option value="2.0" >(GMT +2:00) Kaliningrad, South Africa</option>
                <option value="3.0" >(GMT +3:00) Baghdad, Riyadh, Moscow, St. Petersburg</option>
                <option value="3.5" >(GMT +3:30) Tehran</option>
                <option value="4.0" >(GMT +4:00) Abu Dhabi, Muscat, Baku, Tbilisi</option>
                <option value="4.5" >(GMT +4:30) Kabul</option>
                <option value="5.0" >(GMT +5:00) Ekaterinburg, Islamabad, Karachi, Tashkent</option>
                <option value="5.5" >(GMT +5:30) Bombay, Calcutta, Madras, New Delhi</option>
                <option value="5.75" >(GMT +5:45) Kathmandu</option>
                <option value="6.0" >(GMT +6:00) Almaty, Dhaka, Colombo</option>
                <option value="7.0" >(GMT +7:00) Bangkok, Hanoi, Jakarta</option>
                <option value="8.0" >(GMT +8:00) Beijing, Perth, Singapore, Hong Kong</option>
                <option value="9.0" >(GMT +9:00) Tokyo, Seoul, Osaka, Sapporo, Yakutsk</option>
                <option value="9.5" >(GMT +9:30) Adelaide, Darwin</option>
                <option value="10.0" >(GMT +10:00) Eastern Australia, Guam, Vladivostok</option>
                <option value="11.0" >(GMT +11:00) Magadan, Solomon Islands, New Caledonia</option>
                <option value="12.0" >(GMT +12:00) Auckland, Wellington, Fiji, Kamchatka</option>
            </select>
        </div>

        <button type="submit" class="btn btn-primary btn-block">Save Settings</button>

        <div class="form-group row mt-3">
            <label for="api_key" class="col-sm-3 col-form-label">API Key</label>
            <div class="col-sm-9">
                <input type="text" class="form-control select-input" value="9e657102489b63946908a084befc187e6e506eb0" id="api_key" readonly>
                <small class="form-text text-muted">API Key can be used for read only routes</small>
            </div>
        </div>

        <div class="form-group row">
            <label for="api_secret" class="col-sm-3 col-form-label">API Secret</label>
            <div class="col-sm-9">
                <input type="text" class="form-control select-input" value="6b05b48f4b3a1460f3864c31b26cab6a27dbaff9" id="api_secret" readonly>
                <small class="form-text text-muted">API Secret is used for read, create, update and delete routes</small>
                <small class="form-text text-muted">You can <a class="confirm_btn" data-msg="Are you sure you want to reset the API keys?" href="api/renew">Regenerate API Keys</a> if you need to.</small>
            </div>
        </div>

    </form>
</template>

<script>
import time from '../components/Time'

export default {
  name: 'CoreSettings',
    data () {
        return {
            core: this.$store.getters.core,
        }
    },
    computed: {
      now () {
          return time.now()
      }
    },
    methods: {
        saveSettings () {

        }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
