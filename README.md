This custom command will automatically verify trustworthy users and will automatically perform an action of your choice upon untrustworthy users.
It is based on four key algorithms: No non-ASCII, No numbers, Custom Avatar, and Old Account. It is not fullproof, and as anything else, will have false-positives. It is, however, very accurate.
Make sure that you configure all configurable values to suit your server's needs within the verification.go file.

The steps below are designed to help you save space in your Join message. If you do not need the space, feel free to just copy-paste the code under "Join message in server channel" under Notifications & Feeds > General. And again, configure ALL configurable values.

How to save space with execCC:
Step 1: Navigate to your YagPDB dashboard at https://yagpdb.xyz/manage.
Step 2: Choose your server, then click Home, and under Custom Commands, click "Change settings here."
Step 3: Click "Create a new custom command," set the Trigger Type to "None."
Step 4: Paste the contents of verification.go into the body section (the only other area you can type in), and CONFIGURE THE CONFIGURABLE VALUES TO YOUR SERVER'S NEEDS.
Step 5: Locate the CCID, which will be displayed as #(number) at the top of the page. Keep this number saved.
Step 6: On the sidebar, click Notifications & Feeds, then select General.
Step 7: Enable "Join message in server channel," select any channel (it will not send messages there) and type the following in the body: {{execCC TheCCIDfromStep5 nil 0 .ExecData}}
