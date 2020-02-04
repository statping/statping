<template>
    <div class="alert alert-primary" role="alert">
        <h3>{{message.title}}</h3>
        <span class="mb-3">{{message.description}}</span>
        <div class="d-block mt-2 mb-4">
            <span class="float-left small">
               Started {{toLocal(message.start_on)}} ({{duration(new Date(), message.start_on)}} ago)
            </span>
            <span class="float-right small">
                Ends on {{toLocal(message.end_on)}} (in {{duration(message.end_on, new Date())}})</span>
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
