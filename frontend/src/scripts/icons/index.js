import Vue from "vue";

import Plus from "vue-material-design-icons/Plus.vue";
import Minus from "vue-material-design-icons/Minus.vue";
import ContentSaveOutline from "vue-material-design-icons/ContentSaveOutline.vue";
import FolderOpenOutline from "vue-material-design-icons/FolderOpenOutline.vue";
import ImageFilterCenterFocus from "vue-material-design-icons/ImageFilterCenterFocus.vue";
import TrashCanOutline from "vue-material-design-icons/TrashCanOutline.vue";

import Loading from "vue-material-design-icons/Loading.vue";
import CheckboxMarkedCircleOutline from "vue-material-design-icons/CheckboxMarkedCircleOutline.vue";
import ProgressClose from "vue-material-design-icons/ProgressClose.vue";

Vue.component("icon-save", ContentSaveOutline);
Vue.component("icon-load", FolderOpenOutline);
Vue.component("icon-delete", TrashCanOutline);
Vue.component("icon-zoomin", Plus);
Vue.component("icon-zoomout", Minus);
Vue.component("icon-focus", ImageFilterCenterFocus);

Vue.component("icon-loading", Loading);
Vue.component("icon-ok", CheckboxMarkedCircleOutline);
Vue.component("icon-err", ProgressClose);
