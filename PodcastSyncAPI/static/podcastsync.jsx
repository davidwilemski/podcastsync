/** @jsx React.DOM */
var DownloadPodcastFile = React.createClass({
    render: function() {
        return (
            <div className="download-file">
                <form className="download-form" onSubmit={this.handleSubmit}>
                    <input type="text" id="file-url"></input>
                    <button id="submit-file-download" type="submit" className="btn btn-primary">Send Podcast Episode To Dropbox!</button>
                </form>
            </div>
        );
    },

    handleSubmit: function(e) {
        var creds = this.props.dbclient.credentials();

        $.ajax({
          url: "/podcast/download",
          dataType: 'json',
          type: 'POST',
          data: JSON.stringify({UID: creds.uid, AccessToken: creds.token, PodcastURL: $("#file-url").val()}),
          success: function(data) {
              alert("success! " + data)
            // this.setState({data: data});
          }.bind(this)
        });
        e.preventDefault();
        return false;
    }
});
