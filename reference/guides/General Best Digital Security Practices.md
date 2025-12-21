# **General Best Digital Security Practices**

# **Purpose**

The purpose of this document is to serve as a list of general best practices in digital security. It aims to be useful to individuals or organizations without a dedicated IT/digital security team, with typical digital security and operational security needs. It is highly recommended to conduct a more extensive threat model exercise, especially for those with atypical risks.

It provides various strategies and tools to improve the security of both personal and professional digital communication and storage systems. This includes actionable steps, recommendations on technologies to use and avoid, guidelines for the secure use of email and Google Workspace, Signal messaging, and Zoom, as well as advice on securing hardware such as desktops, laptops, and mobile devices. Additionally, it provides guidance on social media usage, personal safety precautions, and resources for additional support. 

**The first and foremost recommendation is to follow all the requirements that your IT organization has outlined for these topics. In addition to these, we recommend the following best practices.** 

# **Immediate actions**

* Delete unused or old social accounts: To minimize your digital footprint and limit potential exposure of personal information, it's important to delete all unused or old social accounts. This includes removing all old posts that may contain outdated or sensitive information.  
* No public disclosure of personally identifiable information: To protect your privacy, please refrain from sharing any personally identifiable information publicly. This includes location details in any publicly posted content. Delete any posts that contain images of your home, family, or anything that can be used to identify you or your location.  
* Two-factor authentication for social accounts: For all your social media accounts, such as Twitter and LinkedIn, it's crucial to enable two-factor authentication. This adds an extra layer of security to your accounts by requiring a second form of verification beyond your password.  
* Two-factor authentication for personal online accounts: For all your personal online accounts, such as banking and email, enable two-factor authentication immediately. This helps prevent unauthorized access to your accounts by requiring a second form of verification beyond your regular login details.  
* In addition to Two-Factor authentication, everyone should use secure login via Passkeys. This feature is supported by Google Workspace and many other platforms.

# **Tools & Services**

## **Recommendation Summary**

1. Use Password Managers and enable their two-factor authentication, such as 1Password and LastPass.  
2. Set Strong, Unique Passwords for Every Online Account.  
3. Enable Multi-Factor Authentication: Use apps like Authy or Google Authenticator.  
4. Use a VPN for secure, private browsing, such as Mullvad.  
5. Install ad blockers to prevent unwanted ads and potential malware.  
6. Regularly Update Your Software: Always use the latest versions.  
7. Use Encrypted Messaging: Signal is highly recommended.  
8. Use Secure Video Conferencing: Zoom with specific security settings.  
9. Ensure Data Backups: iCloud, Google Cloud are recommended.

## **Discouraged and Encouraged Secure Tech**

Due to known security and cross-platform compatibility issues, certain technologies should not be used for project-related communications. Alternatives are provided for each technology.

Everyone should **not** use:

* SMS/text messages (use Signal instead)  
* Slack (use Signal Groups instead, if possible)  
* iMessage (use Signal instead, as iMessage is insecure for SMS conversations)  
* iCloud Drive (use Google Docs with two-factor protection instead)  
* Office 365 (use Google Docs with two-factor protection instead)  
* DropBox (use Google Drive with two-factor protection instead)  
* Google Meet or Google Chat (use secured Zoom or Signal instead)

**Security Rationale:** To minimize potential points of failure, we aim to reduce the number of platforms used for critical functions. Messaging platforms that lack end-to-end encryption (such as SMS) or are available only on Apple platforms (iOS) are discouraged. Additionally, we advise against using legacy Google applications, such as Hangouts Classic or Duo.

Below are the key often-used technology platforms with recommended security standards:

## **Email**

### **Requirements**

* All users must use two-factor authentication on their email accounts.  
* Email accounts should be managed by a professionally managed IT solution. Google Workspace-based professional accounts are acceptable if professionally managed; personal Gmail accounts are not.  
* All accounts must have a strong password/passphrase of 16 or more characters.  
* Users must not use SMS as their second authentication factor, but may use an authentication app, such as [Authy](https://authy.com/), [Google Authenticator](http://google-authenticator.com/), [1Password Authenticator](https://support.1password.com/one-time-passwords/), or [LastPass Authenticator](https://lastpass.com/auth/?utm_source=impact-radius&utm_medium=affiliate&utm_campaign=affiliate-program&irgwc=1&clickid=QTRWGyUzvxyOTT9wUx0Mo38SUki2lWQxyyUiVI0).  
* Google Hangout Classic and Google Duo should not be used, as they are deprecated products. Instead, secured Zoom or Google Meet in G Suite should be used.

### **Recommendations**

* Users should not click on potential phishing links sent via email until they have verified the source and intent of the link through another medium, such as a phone call or Signal.  
* Account logins and passwords should never be sent as plain text in an email. This is to protect against unauthorized access in case the email account is compromised.  
* We recommend that each member use a hardware key as a second authentication factor. For Google accounts, we recommend a [YubiKey](https://www.yubico.com/product/yubikey-5-series/yubikey-5c-nfc/) or [the Titan Security Key (UTF or Universal Two-Factor)](https://cloud.google.com/titan-security-key/) as the second factor.  
* If permitted by their IT provider, users should use an account with [Google’s Advanced Protection](https://landing.google.com/advancedprotection/).

### **Security Rationale**

Email is one of the oldest technologies still used online, and its security has improved over time. However, due to the need for backward compatibility with older email providers and the variety of email clients and providers, email cannot be fully trusted to be secure. While some email providers support end-to-end encryption, others do not. Many email clients do not visually indicate whether you are sending to an email address that supports this encryption option. Therefore, we recommend replacing email communications with Signal, which supports end-to-end encryption by default and uses a single client per platform when possible.

## **Google Workspace**

Google Docs, Slides, Sheets, Drive, Meets, and Calendar

Using the OFFICIAL VERSION of Google Docs, Slides, and Sheets should be the primary means for document creation and collaboration.

### **Requirements**

* All Google accounts **must use two-step verification** (see above) and can only send to other users who also use it.  
* Everyone accessing Google Drive (for instance, those accessing landscape documents in the shared Google Drive) **must be using two-factor authentication. They must be on a specific set of approved whitelisted users.**  
* All Google accounts **must use a password/passphrase of 16+ characters**.  
* Restrict Google Apps sharing:  
  * **Always restrict viewing/editing permissions of any Google Doc, Sheet, Slide, etc, to specific individuals only.** If we need to share highly secure documents, we can add Advanced Protection protocols that require 2FA physical keys.  
  * **Under sharing settings, ensure “Anyone with this link” is set to OFF** (only specific people can access) rather than View, Comment, or Edit.  
  * Share files directly with specific email addresses. Avoid sharing using the **General access** setting and by **Target Audiences.**  
  * Share Google Doc URLs with groups via Signal as needed, **never via email or unsecured messaging.**  
  * Additional users to any shared Google document URL **can request access to the files and be manually approved as needed.**  
  * **Google Calendar:** “Automatically add invitations” should be disabled on all apps.  
  * **Google Calendar:** Never follow any links you do not recognize from Calendar invites.  
  * **Google Meet:** For enterprise users of Google Mail, [this is now available from inside Gmail’s interface.](https://www.computerworld.com/article/3538916/googles-meet-video-app-gets-gmail-integration.html)  
    * **Google Meet requirements:** We recommend centralizing on Zoom if possible for video and audio conferences, but if Google Meet is required, then use the following standards:  
      * Two-step verification is required for all users  
      * Recordings of the Meet event **must be off.**  
      * **Do not invite users to the Meeting via a URL that is public**  
      * For each meeting, [create a “nickname” for the Meet room](https://support.google.com/meet/answer/9302870#meet). All users will be required to access the room via the Google Meet main menu using their nickname only. This will ensure users cannot join the Meet before the host.

### **Recommendation**

* As with email requirements, if their IT provider allows it, **users should use [Google’s Advanced Protection](https://landing.google.com/advancedprotection/) for their G Suite account.**

### **Security Rationale**

Google’s business applications are widely used, highly scalable, reliable, cross-platform compatible, and generally very secure for most users, with the main security weakness (as with most security systems) being users themselves. Maintaining a requirement for 2FA and being disciplined about how collaborative documents are shared will reduce most of the risk.

## **Signal**

### **Requirements**

* Users **must always use the latest version of Signal** for mobile and desktop, and updates should be installed as new releases are offered.  
* Each member must have the “Use Registration Lock” setting to prevent unwanted new devices from accessing your account.  
* Each member needs to **turn off iCloud call history sharing.** Signal lets you view your call history in the phone app. A useful feature, however, is that if the users have any cloud-based backup system on their phone and computer, they need to turn off this feature in Signal and choose “Show Calls in Recents” to have it Disabled.  
* Users **should have Signal’s Screen Lock settings enabled.**  
* Recommended: Set **“Disappearing Messages” to 1 week.** This gives all contacts time to interact, but after one week, it will automatically remove all posts from all users’ phones.

### **Recommendations**

* **Use Signal for project communication as much as possible.**  
* Recommendation: Signal Video and Audio calls can be used for particularly sensitive conversations to reduce the threat of “man in the middle” eavesdropping. This assumes that our “Show Calls in Recents” setting is off and that the call metadata is secured to the application.  
* Any member who wishes to use Signal but does not want their mobile phone number visible to others can use a service such as Google Voice (assuming two-factor authentication) to register another number in addition to **their primary mobile number.**

### **Security Rationale**

Signal has several security features that make it the best currently available option for secure communication, including end-to-end encryption, a single client per platform, and the option to automatically remove messages from the inbox after a certain amount of time.

Therefore, Signal should **be the primary messaging solution for all text-based IM and group communication.** We recommend using email or SMS only when Signal messaging cannot be used for logistical or technical reasons.

All outreach and communication should be done over Signal by default, unless there is a technical reason why it cannot be.

## **Zoom**

While Zoom has had several notable security issues during the pandemic, they have addressed them with rapid new release versions. Given its widespread use, Zoom chat is a primary choice, but it must be secured in accordance with the following requirements.

### **Requirements**

* **Everyone must use the latest version of the Zoom client on a computer or mobile phone and continue to update to the latest release**.

* Everyone **must use a 16+ digit strong password or passphrase** when using the Zoom application and must use their two-factor-secured email address for login credentials.

* **Links to upcoming Zoom events should not be shared publicly or in emails.** Instead, they can be shared via secure calendar invites.

* Everyone **must use unique, randomly generated passwords for Zoom meetings and meeting IDs for each meeting,** and not share them publicly.

   Meeting passwords should be set so they do not automatically embed inside the Zoom chat URL.

* All Zoom meetings should **not allow attendees to enter before the host.**

* All Zoom meetings **must be set to use only North American data centers.**

* All Zoom meetings **must require all participants to be authenticated to join**.

* Unless needed, all Zoom meetings **should have file sharing disabled.**

* All Zoom meetings should **only allow the host or cohosts to screen share.**

* All Zoom meetings should be set NOT to allow participants to record the meeting, and the autosave chat feature should be OFF.

  * Recording a Zoom meeting also saves any chat messages, even private ones.

### **Recommendations**

* Highly sensitive meetings can be locked after all participants have joined.  
* Highly sensitive meetings can have a Waiting room where each authenticated user is added to the chat individually before the meeting begins.

### **Security Rationale**

During the COVID-19 lockdowns, Zoom became an essential business tool, and though it had some well-publicized security issues as it gained wider use, the bulk of those issues appear to be resolved. Zoom’s scalability, ease of use, and cross-platform compatibility, particularly for calls with large numbers of participants and screen sharing, continue to make it the best available solution for SVF audio/video conferencing and webinar needs. The security measures we recommend align with those of security experts for ensuring Zoom calls are secure.

# **Hardware & Devices**

## **Desktops and Laptops**

### **Recommendations**

1. **Update Operating Systems Regularly:** Ensure you use the latest operating system on your desktops and laptops, regardless of whether you use macOS or Windows. Regularly update all patches and software updates.  
2. **Update All Major Software Applications:** Ensure all major software applications, including web browsers, PDF viewers, and Office applications, are updated to the latest, most secure versions.  
3. **Ensure HTTPS is set by Default:** All major browsers support HTTPS-only mode natively. Confirm that *Always use Secure Connections* is enforced.  
4. **Remove Adobe Flash:** Disable Adobe Flash completely or set it to “click to run” on your desktop or laptop web browser. This plugin is known for its security vulnerabilities.  
5. **Set WIFI to Use WPA2 Password Encryption:** For both home and work WIFI, ensure that you are set to use the more secure WPA2 password encryption.  
6. **Backup Data in Real Time:** Ensure your laptop, desktop, and network content are backed up in real time or near real time. iCloud, Google Cloud, or other similar platforms are highly recommended for personal devices.  
7. **Full Disk Encryption:** Utilize full disk encryption on laptops, either using BitLocker (Windows) or FileVault (Mac). Alternatively, use VeraCrypt on Windows devices.

### **Security Rationale**

Desktops and laptops are critical points of vulnerability when they lack the most recent security patches and updates, including essential software applications such as browsers, web plugins, Office Applications, and networks. Enhancing the security of these devices significantly involves regularly updating operating systems and software applications, using encrypted connections, removing insecure plugins, enabling secure Wi-Fi password encryption, and utilizing real-time backup services.

## **Mobile Devices**

### **Recommendations**

1. **Use Recent Versions of Mobile Operating Systems:** We recommend all users use a recent version of their smartphone's operating system. For Apple iPhone users, ensure you have the latest iOS version and install new software updates as they become available. Android users should also ensure they are using the latest version of Android.  
2. **Secure Login and Password:** For Apple users, use complex passwords that meet our recommended standards for your Apple ID. Also, use Two-Factor Authentication for your Apple ID. For Android users, ensure your Google account password is strong and enable Two-Factor Authentication.  
3. **Manage Location Services:** For both iOS and Android users, change the settings so that only very critical applications (such as Maps applications) have access to your "location services." Even then, only set these applications to access location services while they are running.  
4. **Encrypt Backups:** For iPhone users, ensure that iTunes and iCloud encrypt all data backups. Android users should use Google's backup service, which automatically encrypts backups.  
5. **Limit Logins:** For both iPhone and Android users, set your device to automatically wipe all data after a certain number of failed passcode attempts.

### **Security Rationale**

In today's world, we increasingly work and live through our mobile devices, which carry key communications, emails, passwords, and other data. This makes them vulnerable to attacks, especially voice communication, internet chats, and call information. The Android system is somewhat less secure than the Apple iPhone platform, as it is more open to security breaches. Therefore, we strongly recommend a series of actions to enhance the security of your mobile devices. These include regularly updating operating systems, using encrypted connections, limiting login attempts, encrypting backups, and managing location services. By implementing these measures, we can significantly improve the security of our mobile devices.

# **Social Media**

## **Recommendations**

* Delete Old Posts: Remove any outdated or irrelevant posts from your social media platforms. This can help reduce the risk of malicious use of unwanted information.  
* Audit Your Own Accounts: Regularly review your social media accounts to ensure all information is accurate and up to date. This includes checking privacy settings, reviewing friend lists, and deleting any unnecessary information.  
* Make Accounts Private: Adjust your social media settings so that only friends or approved followers can see your posts. This can greatly reduce the risk of unwanted individuals or entities accessing your information.  
* Separate Work from Personal Accounts: Try to keep your professional and personal social media accounts separate. This can help protect your personal information from exposure during your professional interactions.  
* Remove Family Details and Information: To protect your family members' privacy and safety, avoid sharing sensitive details about them on social media. This includes their full names, locations, schools, or workplaces.  
* Avoid Using Geotag Functions on Social Media Posts: Geotagging can reveal your exact location to anyone who views your post. To protect your privacy, it's best to avoid using this function.  
* Avoid Posting Photos That Include Identifying Information: Try not to share pictures that include things like house numbers, car plates, or anything else that could be used to identify you or your location.  
* Use Privacy Tools: Consider using online tools, such as Block Party (https://www.blockpartyapp.com/privacy-party/), to help you control what you see and who can interact with you on social media platforms. This can help reduce the risk of harassment or unwanted interactions.

# **Personal safety**

## **Recommendations**

* Google Alerts on Who You Are: Set up Google Alerts for your name and other personal identifiers to monitor the internet for any new mentions or uses of your information.  
* PimEye Yourself: Use the PimEye facial recognition search engine to see where images of your face may be posted online. This can help you identify any unauthorized or unwanted uses of your photos.  
* Turn on Credit Freezes: Activate credit freezes with all major credit reporting agencies to prevent unauthorized individuals from learning detailed personal information about you or opening new accounts in your name.  
* Have a Family “Code Word” to Verify Yourself: Establish a code word or phrase with your family members that can be used to verify identity in communications, especially in situations that might involve impersonation or fraud.  
* Transition Family to Signal: Encourage your family members to start using Signal, a highly secure messaging app, for all family communications to enhance privacy and security.  
* Invest in a Privacy Service: Consider subscribing to a privacy service that provides additional protections, such as encrypted email, virtual private networks (VPNs), and data breach alerts.  
* Invest in a Data Broker Removal Service: Consider subscribing to a removal service for you and your family to make doxxing you or your family more difficult.  
* Check on "Have I Been Pwned": Regularly enter your email addresses into the "Have I Been Pwned" website to check if your accounts have been compromised in any recent data breaches.

# **Getting help**

Know who to turn to

* If you have one, your own IT department  
* The network’s helpdesk, if it exists

# **Checklist**

The following is a checklist derived from the document for an IT consultant to use during self-assessments:

### **Security Setup and Practices**

- [ ] **Enable Two-Factor Authentication (2FA)**: Verify if 2FA is enabled on all accounts, using tools like 1Password or LastPass.  
- [ ] **Set Strong, Unique Passwords**: Ensure all online accounts use strong, unique passwords.  
- [ ] **Enable Multi-Factor Authentication (MFA)**: Ensure MFA is activated wherever possible, using apps like Authy or Google Authenticator.  
- [ ] **Use a VPN**: Ensure a VPN is in use for secure, private browsing.  
- [ ] **Install Ad-Blockers**: Ensure ad blockers are installed to prevent unwanted ads and potential malware.  
- [ ] **Software Updates**: Ensure that all software is regularly updated to the latest versions.  
- [ ] **Encrypted Messaging**: Verify the use of encrypted messaging apps like Signal.  
- [ ] **Secure Video Conferencing**: Check if Zoom is being used with the recommended security settings.  
- [ ] **Data Backups**: Ensure regular backups are in place using services such as iCloud or Google Cloud.

### **Technology Use: Discouraged and Encouraged**

- [ ] **Avoid Certain Technologies**: Confirm that users are not using discouraged technologies (e.g., SMS/text messages, Slack, iMessage, iCloud Drive, Office 365, Dropbox, Google Meet/Chat) for project-related communications.  
- [ ] **Encouraged Technologies**: Ensure the use of recommended secure alternatives (e.g., Signal for messaging, Google Docs with two-factor protection, secured Zoom).

      ### **Email Security**

- [ ] **Two-Factor Authentication for Email**: Verify that 2FA is enabled on all email accounts.  
- [ ] **Professionally Managed IT for Email**: Ensure email accounts are managed by a professional IT solution, preferably Google Workspace-based accounts.

### **Social Media and Personal Information Security**

- [ ] **Monitor Personal Information**: Check if Google Alerts are set up for personal identifiers to monitor the internet for new mentions.  
- [ ] **Online Presence**: Use PimEye or similar services to check where images of faces are posted online.  
- [ ] **Credit Freezes**: Confirm that credit freezes are activated with all major credit reporting agencies.  
- [ ] **Family Security Practices**: Ensure a family “code word” is established, transition family communications to Signal, and consider investing in a privacy service.

### **General Security Measures**

- [ ] **Regular Checks on "Have I Been Pwned"**: Regularly check email addresses on "Have I Been Pwned" to see if accounts have been compromised in data breaches.  
- [ ] **IT Support Awareness**: Ensure users know who to contact for IT support, including a network's helpdesk or IT department, if available.

