<template>
  <ul class="pagination d-flex justify-content-center">
    <li class="page-item">
      <button
        class="btn btn-outline-secondary page-link"
        aria-label="Previous"
        :disabled="skip <= 0"
        @click="getPrevDowntimes"
      >
        <span aria-hidden="true">
          &laquo;
        </span>
      </button>
    </li>
  
    <li class="page-item">
      <button
        class="btn btn-outline-secondary page-link"
        aria-label="Next"
        @click="getNextDowntimes"
      >
        <span aria-hidden="true">
          &raquo;
        </span>
      </button>
    </li>
  </ul>
</template>

<script>
export default {
    props: {
        getNextDowntimes: {
            type: Function,
            default: function () {}
        },
        getPrevDowntimes: {
            type: Function,
            default: function () {}
        },
        skip: {
            type: Number,
            default: 0,
        }
    }
};
</script>