/** @jsx React.DOM */
var DownloadPodcastFile = React.createClass({
  getInitialState: function() {
    return {url: ""};
  },
    render: function() {
        return (
            <div className="download-file">
                <form className="download-form" onSubmit={this.handleSubmit}>
                    <input type="text" id="file-url" value={this.state.url} onChange={this.handleInput}></input>
                    <button id="submit-file-download" type="submit" className="btn btn-primary">Send Podcast Episode To Dropbox!</button>
                </form>
            </div>
        );
    },
    handleSubmit: function(e) {
      var creds = this.props.dbclient.credentials();

      $.ajax({
        url: "/podcast/download",
        type: 'POST',
        data: JSON.stringify({UID: creds.uid, AccessToken: creds.token, PodcastURL: $("#file-url").val()}),
      }).done(function(data, textStatus) {
        console.log(data)
        console.log(textStatus);
        window.alert("success! " + data);
        this.setState({url: ""});
      }.bind(this)).fail(function( jqXHR, textStatus, errorThrown ) {
        console.log(textStatus);
        console.log(errorThrown);
        window.alert("Error: invalid URL");
      });

      return false;
    },
    handleInput: function(e) {
      this.setState({url: e.target.value})
    }
});
