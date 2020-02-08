<template>
    <div class="col-12 full-col-12">
        <h4 v-if="group.name !== 'Empty Group'" class="group_header">{{group.name}}</h4>
        <div class="list-group online_list mb-3">

            <a v-for="(service, index) in $store.getters.servicesInGroup(group.id)" v-bind:key="index" class="service_li list-group-item list-group-item-action">
                {{service.name}}
                <span class="badge bg-success float-right pulse-glow">{{service.online ? "ONLINE" : "OFFLINE"}}</span>
            </a>

        </div>
    </div>
</template>

<script>

export default {
  name: 'Group',
  components: {

  },
  props: {
    group: Object
  },
    methods: {
      serviceBadge (s) {

      },
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    @keyframes pulse_animation {
        0% { transform: scale(1); }
        30% { transform: scale(1); }
        40% { transform: scale(1.02); }
        50% { transform: scale(1); }
        60% { transform: scale(1); }
        70% { transform: scale(1.05); }
        80% { transform: scale(1); }
        100% { transform: scale(1); }
    }

    .pulse {
        animation-name: pulse_animation;
        animation-duration: 1500ms;
        transform-origin:70% 70%;
        animation-iteration-count: infinite;
        animation-timing-function: linear;
    }


    @keyframes glow-grow {
        0% {
            opacity: 0;
            transform: scale(1);
        }
        80% {
            opacity: 1;
        }
        100% {
            transform: scale(2);
            opacity: 0;
        }
    }
    .pulse-glow {
        animation-name: glow-grown;
        animation-duration: 100ms;
        transform-origin: 70% 30%;
        animation-iteration-count: infinite;
        animation-timing-function: linear;
    }

    .pulse-glow:before,
    .pulse-glow:after {
        position: absolute;
        content: "";
        height: 0.4rem;
        width: 1.7rem;
        top: 1.3rem;
        right: 2.15rem;
        border-radius: 0;
        box-shadow: 0 0 6px #47d337;
        animation: glow-grow 2s ease-out infinite;
    }
</style>
