#ifndef __RTC_CLIENT_H__
#define __RTC_CLIENT_H__

#include <string>
#include "rtc_manager.h"
#include "rtc_connection.h"
#include "RtcWinUtils.h"

class RtcClientStateObserver {
public:
	virtual void OnIceCandidate(std::string uid, const std::string mid, const std::string sdp) = 0;
};

class RtcClient : public RTCMessageSender {
public:
	RtcClient(std::string uid, std::string name, std::shared_ptr<RTCManager> manager, bool local) {
		is_local_ = local;
		user_name_ = name;
		user_id_ = uid;
		rtc_manager_ = manager;
		rtc_connection_ = nullptr;
		observer_ = nullptr;
	}

	~RtcClient() {
		DestroyPeerConnection();
	}

	bool CreatePeerConnection() {
		if (rtc_manager_ && rtc_connection_ == nullptr) {
			webrtc::PeerConnectionInterface::RTCConfiguration rtc_config;

			rtc_connection_ = rtc_manager_->CreateConnection(rtc_config, this);
			if (is_local_ && rtc_connection_) {
				rtc_manager_->CreateLocalMediaStream(rtc_connection_);
			}
		}
		return rtc_connection_ != nullptr;
	}

	void DestroyPeerConnection() {
		if (is_local_) {
			rtc_manager_->DestroyLocalMediaStream(rtc_connection_);
		}
		rtc_connection_ = nullptr;
	}

	std::string GetUserName() {
		return user_name_;
	}

	std::string GetUserId() {
		return user_id_;
	}

	bool IsLocalUser() {
		return is_local_;
	}

	std::shared_ptr<RTCConnection> GetConnection() {
		return rtc_connection_;
	}

	virtual void OnIceConnectionStateChange(webrtc::PeerConnectionInterface::IceConnectionState new_state) {
		if (new_state >= 4) {
			LogPrintf("Connection Failed!!!");
		}
	}

	virtual void OnIceCandidate(const std::string sdp_mid, const int sdp_mlineindex, const std::string sdp) {
		if (observer_) {
			observer_->OnIceCandidate(user_id_, sdp_mid, sdp);
		}
	}

	void SetStateObserver(RtcClientStateObserver* observer) {
		observer_ = observer;
	}

private:
	bool        is_local_;
	std::string user_name_;
	std::string user_id_;

	std::shared_ptr<RTCManager> rtc_manager_;
	std::shared_ptr<RTCConnection> rtc_connection_;
	RtcClientStateObserver* observer_;
};


#endif