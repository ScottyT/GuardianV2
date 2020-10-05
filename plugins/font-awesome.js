import Vue from 'vue';
import {
  library,
  config
} from '@fortawesome/fontawesome-svg-core';
import {
  FontAwesomeIcon, FontAwesomeLayers
} from '@fortawesome/vue-fontawesome'
import {
  faPlus,
  faCircle
} from '@fortawesome/free-solid-svg-icons';
import {
  faFacebookSquare, faLinkedin
} from '@fortawesome/free-brands-svg-icons'
config.autoAddCss = true;
library.add(faCircle, faPlus, faFacebookSquare, faLinkedin);
Vue.component('font-awesome-icon', FontAwesomeIcon);
Vue.component('font-awesome-layers', FontAwesomeLayers);