<!--
  - Statup
  - Copyright (C) 2020.  Hunter Long and the project contributors
  - Written by Hunter Long <info@socialeck.com> and the project contributors
  -
  - https://github.com/hunterlong/statup
  -
  - The licenses for most software and other practical works are designed
  - to take away your freedom to share and change the works.  By contrast,
  - the GNU General Public License is intended to guarantee your freedom to
  - share and change all versions of a program--to make sure it remains free
  - software for all its users.
  -
  - You should have received a copy of the GNU General Public License
  - along with this program.  If not, see <http://www.gnu.org/licenses/>.
  -->

<template>
    <div>
    <div class="col-12">
        <h1 class="text-black-50">Messages</h1>
        <table class="table table-striped">
            <thead>
            <tr>
                <th scope="col">Title</th>
                <th scope="col" class="d-none d-md-table-cell">Service</th>
                <th scope="col" class="d-none d-md-table-cell">Begins</th>
                <th scope="col"></th>
            </tr>
            </thead>
            <tbody>

            <tr v-for="(message, index) in messages" v-bind:key="index">
                <td>{{message.title}}</td>
                <td class="d-none d-md-table-cell"><a href="service/1">{{message.service}}</a></td>
                <td class="d-none d-md-table-cell">{{message.start_on}}</td>
                <td class="text-right">
                    <div class="btn-group">
                        <a href="message/1" class="btn btn-outline-secondary"><i class="fas fa-exclamation-triangle"></i> Edit</a>
                        <a href="api/messages/1" class="ajax_delete btn btn-danger"><i class="fas fa-times"></i></a>
                    </div>
                </td>
            </tr>

            </tbody>
        </table>
    </div>


    <div class="col-12">
        <h1 class="text-black-50 mt-5">Create Message</h1>

        <div class="card">
            <div class="card-body">

                <form class="ajax_form" action="api/messages" data-redirect="messages" method="POST">
                    <div class="form-group row">
                        <label for="username" class="col-sm-4 col-form-label">Title</label>
                        <div class="col-sm-8">
                            <input type="text" name="title" class="form-control" value="" id="title" placeholder="Message Title" required>
                        </div>
                    </div>

                    <div class="form-group row">
                        <label for="username" class="col-sm-4 col-form-label">Description</label>
                        <div class="col-sm-8">
                            <textarea rows="5" name="description" class="form-control" id="description" required></textarea>
                        </div>
                    </div>

                    <div class="form-group row">
                        <label class="col-sm-4 col-form-label">Message Date Range</label>
                        <div class="col-sm-4">
                            <input type="text" name="start_on" class="form-control form-control-plaintext" id="start_on" value="0001-01-01T00:00:00Z" required>
                        </div>
                        <div class="col-sm-4">
                            <input type="text" name="end_on" class="form-control form-control-plaintext" id="end_on" value="0001-01-01T00:00:00Z" required>
                        </div>
                    </div>

                    <div class="form-group row">
                        <label for="service_id" class="col-sm-4 col-form-label">Service</label>
                        <div class="col-sm-8">
                            <select class="form-control" name="service" id="service_id">
                                <option value="0" selected>Global Message</option>


                                <option value="7" >Statping API</option>


                                <option value="6" >Push Notification Server</option>


                                <option value="1" >Google</option>


                                <option value="2" >Statping Github</option>


                                <option value="3" >JSON Users Test</option>


                                <option value="4" >JSON API Tester</option>


                                <option value="5" >Google DNS</option>

                            </select>
                        </div>
                    </div>

                    <div class="form-group row">
                        <label for="notify_method" class="col-sm-4 col-form-label">Notification Method</label>
                        <div class="col-sm-8">
                            <input type="text" name="notify_method" class="form-control" id="notify_method" value="" placeholder="email">
                        </div>
                    </div>

                    <div class="form-group row">
                        <label for="notify_method" class="col-sm-4 col-form-label">Notify Users</label>
                        <div class="col-sm-8">
            <span class="switch">
            <input type="checkbox" name="notify_users-value" class="switch" id="switch-normal">
            <label for="switch-normal">Notify Users Before Scheduled Time</label>
            <input type="hidden" name="notify_users" id="switch-normal-value" value="false">
          </span>
                        </div>
                    </div>

                    <div class="form-group row">
                        <label for="notify_before" class="col-sm-4 col-form-label">Notify Before</label>
                        <div class="col-sm-8">
                            <div class="form-inline">
                                <input type="number" name="notify_before" class="col-4 form-control" id="notify_before" value="0">
                                <select class="ml-2 col-7 form-control" name="notify_before_scale" id="notify_before_scale">
                                    <option value="minute">Minutes</option>
                                    <option value="hour">Hours</option>
                                    <option value="day">Days</option>
                                </select>
                            </div>
                        </div>
                    </div>

                    <div class="form-group row">
                        <div class="col-sm-12">
                            <button type="submit" class="btn btn-primary btn-block">Create Message</button>
                        </div>
                    </div>
                    <div class="alert alert-danger d-none" id="alerter" role="alert"></div>
                </form>
            </div>
        </div>

    </div>
    </div>
</template>

<script>
  import Api from "../API"

  export default {
  name: 'DashboardMessages',
  data () {
    return {
        messages: null
    }
  },
  created() {
    this.getMessages()
  },
  methods: {
    async getMessages () {
      this.messages = await Api.messages()
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
