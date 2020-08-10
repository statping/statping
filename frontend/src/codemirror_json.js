import CodeMirror from 'codemirror'

CodeMirror.defineMode('mymode', () => {
  return {
    token(stream, state) {
      if (stream.match(".Service") || (stream.match(".Core")) || (stream.match(".Failure"))) {
        return "var-highlight"
      } else if (stream.match(".Id") || stream.match(".Domain") || stream.match(".CreatedAt") ||
        stream.match(".Name") || stream.match(".Downtime.Human") || stream.match(".Issue") || stream.match(".LastStatusCode") ||
        stream.match(".Port") || stream.match(".FailuresLast24Hours") || stream.match(".PingTime")) {
        return "var-sub-highlight"
      } else if (stream.match("{{") || stream.match("}}")) {
        return "bracketer"
      } else {
        stream.next()
        return null
      }
    }
  }
})
