// --credential-file

package twitter_v2

func flags() {

}

const search_tweets_py_h = `
    usage: search_tweets.py [-h] [--credential-file CREDENTIAL_FILE]
                      [--credential-file-key CREDENTIAL_YAML_KEY]
                      [--env-overwrite ENV_OVERWRITE]
                      [--config-file CONFIG_FILENAME]
                      [--account-type {premium,enterprise}]
                      [--count-bucket COUNT_BUCKET]
                      [--start-datetime FROM_DATE] [--end-datetime TO_DATE]
                      [--filter-rule PT_RULE]
                      [--results-per-call RESULTS_PER_CALL]
                      [--max-results MAX_RESULTS] [--max-pages MAX_PAGES]
                      [--results-per-file RESULTS_PER_FILE]
                      [--filename-prefix FILENAME_PREFIX]
                      [--no-print-stream] [--print-stream]
                      [--extra-headers EXTRA_HEADERS] [--debug]

optional arguments:
  -h, --help            show this help message and exit
  --credential-file CREDENTIAL_FILE
                        Location of the yaml file used to hold your
                        credentials.
  --credential-file-key CREDENTIAL_YAML_KEY
                        the key in the credential file used for this session's
                        credentials. Defaults to search_tweets_api
  --env-overwrite ENV_OVERWRITE
                        Overwrite YAML-parsed credentials with any set
                        environment variables. See API docs or readme for
                        details.
  --config-file CONFIG_FILENAME
                        configuration file with all parameters. Far, easier to
                        use than the command-line args version., If a valid
                        file is found, all args will be populated, from there.
                        Remaining command-line args, will overrule args found
                        in the config, file.
  --account-type {premium,enterprise}
                        The account type you are using
  --count-bucket COUNT_BUCKET
                        Set this to make a 'counts' request. Bucket size for counts endpoint. Options:, day, hour,
                        minute.
  --start-datetime FROM_DATE
                        Start of datetime window, format 'YYYY-mm-DDTHH:MM'
                        (default: -30 days)
  --end-datetime TO_DATE
                        End of datetime window, format 'YYYY-mm-DDTHH:MM'
                        (default: most recent date)
  --filter-rule PT_RULE
                        PowerTrack filter rule (See: http://support.gnip.com/c
                        ustomer/portal/articles/901152-powertrack-operators)
  --results-per-call RESULTS_PER_CALL
                        Number of results to return per call (default 100; max
                        500) - corresponds to 'maxResults' in the API. If making a 'counts' request with '--count-bucket', this parameter is ignored.
  --max-results MAX_RESULTS
                        Maximum number of Tweets or Counts to return for this
                        session (defaults to 500)
  --max-pages MAX_PAGES
                        Maximum number of pages/API calls to use for this
                        session.
  --results-per-file RESULTS_PER_FILE
                        Maximum tweets to save per file.
  --filename-prefix FILENAME_PREFIX
                        prefix for the filename where tweet json data will be
                        stored.
  --no-print-stream     disable print streaming
  --print-stream        Print tweet stream to stdout
  --extra-headers EXTRA_HEADERS
                        JSON-formatted str representing a dict of additional
                        request headers
  --debug               print all info and warning messages`
