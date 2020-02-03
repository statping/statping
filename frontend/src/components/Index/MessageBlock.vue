<template>
    <div class="alert alert-primary" role="alert">
        <h3>{{message.title}}</h3>
        <span class="mb-3">{{message.description}}</span>
        <div class="d-block mt-2 mb-4">
            <span class="float-left small">
               Starts on {{parseTime(message.start_on)}}
            </span>
            <span class="float-right small">
                Ends on {{parseTime(message.end_on)}}</span>
        </div>
    </div>
</template>

<script>
export default {
  name: 'MessageBlock',
    props: {
        message: {
            type: Object
        }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
