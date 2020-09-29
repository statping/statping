import {library} from '@fortawesome/fontawesome-svg-core'
import {fas} from '@fortawesome/fontawesome-free-solid';
import {fab} from '@fortawesome/free-brands-svg-icons';
import {FontAwesomeIcon} from '@fortawesome/vue-fontawesome'
import Vue from "vue";

library.add(fas, fab)

Vue.component('font-awesome-icon', FontAwesomeIcon)
