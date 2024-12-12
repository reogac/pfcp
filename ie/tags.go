package ie

/* PFCP message type */
const (
	MSG_PFCP_HEARTBEAT_REQUEST_TYPE                 uint32 = 1
	MSG_PFCP_HEARTBEAT_RESPONSE_TYPE                uint32 = 2
	MSG_PFCP_PFD_MANAGEMENT_REQUEST_TYPE            uint32 = 3
	MSG_PFCP_PFD_MANAGEMENT_RESPONSE_TYPE           uint32 = 4
	MSG_PFCP_ASSOCIATION_SETUP_REQUEST_TYPE         uint32 = 5
	MSG_PFCP_ASSOCIATION_SETUP_RESPONSE_TYPE        uint32 = 6
	MSG_PFCP_ASSOCIATION_UPDATE_REQUEST_TYPE        uint32 = 7
	MSG_PFCP_ASSOCIATION_UPDATE_RESPONSE_TYPE       uint32 = 8
	MSG_PFCP_ASSOCIATION_RELEASE_REQUEST_TYPE       uint32 = 9
	MSG_PFCP_ASSOCIATION_RELEASE_RESPONSE_TYPE      uint32 = 10
	MSG_PFCP_VERSION_NOT_SUPPORTED_RESPONSE_TYPE    uint32 = 11
	MSG_PFCP_NODE_REPORT_REQUEST_TYPE               uint32 = 12
	MSG_PFCP_NODE_REPORT_RESPONSE_TYPE              uint32 = 13
	MSG_PFCP_SESSION_SET_DELETION_REQUEST_TYPE      uint32 = 14
	MSG_PFCP_SESSION_SET_DELETION_RESPONSE_TYPE     uint32 = 15
	MSG_PFCP_SESSION_SET_MODIFICATION_REQUEST_TYPE  uint32 = 16
	MSG_PFCP_SESSION_SET_MODIFICATION_RESPONSE_TYPE uint32 = 17
	MSG_PFCP_SESSION_ESTABLISHMENT_REQUEST_TYPE     uint32 = 50
	MSG_PFCP_SESSION_ESTABLISHMENT_RESPONSE_TYPE    uint32 = 51
	MSG_PFCP_SESSION_MODIFICATION_REQUEST_TYPE      uint32 = 52
	MSG_PFCP_SESSION_MODIFICATION_RESPONSE_TYPE     uint32 = 53
	MSG_PFCP_SESSION_DELETION_REQUEST_TYPE          uint32 = 54
	MSG_PFCP_SESSION_DELETION_RESPONSE_TYPE         uint32 = 55
	MSG_PFCP_SESSION_REPORT_REQUEST_TYPE            uint32 = 56
	MSG_PFCP_SESSION_REPORT_RESPONSE_TYPE           uint32 = 57
)

/* Group/element type */
const (
	CREATE_PDR_TYPE                                                               uint32 = 1
	PDI_TYPE                                                                      uint32 = 2
	CREATE_FAR_TYPE                                                               uint32 = 3
	FORWARDING_PARAMETERS_TYPE                                                    uint32 = 4
	DUPLICATING_PARAMETERS_TYPE                                                   uint32 = 5
	CREATE_URR_TYPE                                                               uint32 = 6
	CREATE_QER_TYPE                                                               uint32 = 7
	CREATED_PDR_TYPE                                                              uint32 = 8
	UPDATE_PDR_TYPE                                                               uint32 = 9
	UPDATE_FAR_TYPE                                                               uint32 = 10
	UPDATE_FORWARDING_PARAMETERS_TYPE                                             uint32 = 11
	UPDATE_BAR_PFCP_SESSION_REPORT_RESPONSE_TYPE                                  uint32 = 12
	UPDATE_URR_TYPE                                                               uint32 = 13
	UPDATE_QER_TYPE                                                               uint32 = 14
	REMOVE_PDR_TYPE                                                               uint32 = 15
	REMOVE_FAR_TYPE                                                               uint32 = 16
	REMOVE_URR_TYPE                                                               uint32 = 17
	REMOVE_QER_TYPE                                                               uint32 = 18
	CAUSE_TYPE                                                                    uint32 = 19
	SOURCE_INTERFACE_TYPE                                                         uint32 = 20
	F_TEID_TYPE                                                                   uint32 = 21
	NETWORK_INSTANCE_TYPE                                                         uint32 = 22
	SDF_FILTER_TYPE                                                               uint32 = 23
	APPLICATION_ID_TYPE                                                           uint32 = 24
	GATE_STATUS_TYPE                                                              uint32 = 25
	MBR_TYPE                                                                      uint32 = 26
	GBR_TYPE                                                                      uint32 = 27
	QER_CORRELATION_ID_TYPE                                                       uint32 = 28
	PRECEDENCE_TYPE                                                               uint32 = 29
	TRANSPORT_LEVEL_MARKING_TYPE                                                  uint32 = 30
	VOLUME_THRESHOLD_TYPE                                                         uint32 = 31
	TIME_THRESHOLD_TYPE                                                           uint32 = 32
	MONITORING_TIME_TYPE                                                          uint32 = 33
	SUBSEQUENT_VOLUME_THRESHOLD_TYPE                                              uint32 = 34
	SUBSEQUENT_TIME_THRESHOLD_TYPE                                                uint32 = 35
	INACTIVITY_DETECTION_TIME_TYPE                                                uint32 = 36
	REPORTING_TRIGGERS_TYPE                                                       uint32 = 37
	REDIRECT_INFORMATION_TYPE                                                     uint32 = 38
	REPORT_TYPE_TYPE                                                              uint32 = 39
	OFFENDING_IE_TYPE                                                             uint32 = 40
	FORWARDING_POLICY_TYPE                                                        uint32 = 41
	DESTINATION_INTERFACE_TYPE                                                    uint32 = 42
	UP_FUNCTION_FEATURES_TYPE                                                     uint32 = 43
	APPLY_ACTION_TYPE                                                             uint32 = 44
	DOWNLINK_DATA_SERVICE_INFORMATION_TYPE                                        uint32 = 45
	DOWNLINK_DATA_NOTIFICATION_DELAY_TYPE                                         uint32 = 46
	DL_BUFFERING_DURATION_TYPE                                                    uint32 = 47
	DL_BUFFERING_SUGGESTED_PACKET_COUNT_TYPE                                      uint32 = 48
	PFCPSMREQ_FLAGS_TYPE                                                          uint32 = 49
	PFCPSRRSP_FLAGS_TYPE                                                          uint32 = 50
	LOAD_CONTROL_INFORMATION_TYPE                                                 uint32 = 51
	SEQUENCE_NUMBER_TYPE                                                          uint32 = 52
	METRIC_TYPE                                                                   uint32 = 53
	OVERLOAD_CONTROL_INFORMATION_TYPE                                             uint32 = 54
	TIMER_TYPE                                                                    uint32 = 55
	PDR_ID_TYPE                                                                   uint32 = 56
	F_SEID_TYPE                                                                   uint32 = 57
	APPLICATION_ID_S_PFDS_TYPE                                                    uint32 = 58
	PFD_CONTEXT_TYPE                                                              uint32 = 59
	NODE_ID_TYPE                                                                  uint32 = 60
	PFD_CONTENTS_TYPE                                                             uint32 = 61
	MEASUREMENT_METHOD_TYPE                                                       uint32 = 62
	USAGE_REPORT_TRIGGER_TYPE                                                     uint32 = 63
	MEASUREMENT_PERIOD_TYPE                                                       uint32 = 64
	FQ_CSID_TYPE                                                                  uint32 = 65
	VOLUME_MEASUREMENT_TYPE                                                       uint32 = 66
	DURATION_MEASUREMENT_TYPE                                                     uint32 = 67
	APPLICATION_DETECTION_INFORMATION_TYPE                                        uint32 = 68
	TIME_OF_FIRST_PACKET_TYPE                                                     uint32 = 69
	TIME_OF_LAST_PACKET_TYPE                                                      uint32 = 70
	QUOTA_HOLDING_TIME_TYPE                                                       uint32 = 71
	DROPPED_DL_TRAFFIC_THRESHOLD_TYPE                                             uint32 = 72
	VOLUME_QUOTA_TYPE                                                             uint32 = 73
	TIME_QUOTA_TYPE                                                               uint32 = 74
	START_TIME_TYPE                                                               uint32 = 75
	END_TIME_TYPE                                                                 uint32 = 76
	QUERY_URR_TYPE                                                                uint32 = 77
	USAGE_REPORT_SESSION_MODIFICATION_RESPONSE_TYPE                               uint32 = 78
	USAGE_REPORT_SESSION_DELETION_RESPONSE_TYPE                                   uint32 = 79
	USAGE_REPORT_SESSION_REPORT_REQUEST_TYPE                                      uint32 = 80
	URR_ID_TYPE                                                                   uint32 = 81
	LINKED_URR_ID_TYPE                                                            uint32 = 82
	DOWNLINK_DATA_REPORT_TYPE                                                     uint32 = 83
	OUTER_HEADER_CREATION_TYPE                                                    uint32 = 84
	CREATE_BAR_TYPE                                                               uint32 = 85
	UPDATE_BAR_SESSION_MODIFICATION_REQUEST_TYPE                                  uint32 = 86
	REMOVE_BAR_TYPE                                                               uint32 = 87
	BAR_ID_TYPE                                                                   uint32 = 88
	CP_FUNCTION_FEATURES_TYPE                                                     uint32 = 89
	USAGE_INFORMATION_TYPE                                                        uint32 = 90
	APPLICATION_INSTANCE_ID_TYPE                                                  uint32 = 91
	FLOW_INFORMATION_TYPE                                                         uint32 = 92
	UE_IP_ADDRESS_TYPE                                                            uint32 = 93
	PACKET_RATE_TYPE                                                              uint32 = 94
	OUTER_HEADER_REMOVAL_TYPE                                                     uint32 = 95
	RECOVERY_TIME_STAMP_TYPE                                                      uint32 = 96
	DL_FLOW_LEVEL_MARKING_TYPE                                                    uint32 = 97
	HEADER_ENRICHMENT_TYPE                                                        uint32 = 98
	ERROR_INDICATION_REPORT_TYPE                                                  uint32 = 99
	MEASUREMENT_INFORMATION_TYPE                                                  uint32 = 100
	NODE_REPORT_TYPE_TYPE                                                         uint32 = 101
	USER_PLANE_PATH_FAILURE_REPORT_TYPE                                           uint32 = 102
	REMOTE_GTP_U_PEER_TYPE                                                        uint32 = 103
	UR_SEQN_TYPE                                                                  uint32 = 104
	UPDATE_DUPLICATING_PARAMETERS_TYPE                                            uint32 = 105
	ACTIVATE_PREDEFINED_RULES_TYPE                                                uint32 = 106
	DEACTIVATE_PREDEFINED_RULES_TYPE                                              uint32 = 107
	FAR_ID_TYPE                                                                   uint32 = 108
	QER_ID_TYPE                                                                   uint32 = 109
	OCI_FLAGS_TYPE                                                                uint32 = 110
	PFCP_ASSOCIATION_RELEASE_REQUEST_TYPE                                         uint32 = 111
	GRACEFUL_RELEASE_PERIOD_TYPE                                                  uint32 = 112
	PDN_TYPE_TYPE                                                                 uint32 = 113
	FAILED_RULE_ID_TYPE                                                           uint32 = 114
	TIME_QUOTA_MECHANISM_TYPE                                                     uint32 = 115
	USER_PLANE_IP_RESOURCE_INFORMATION_TYPE                                       uint32 = 116
	USER_PLANE_INACTIVITY_TIMER_TYPE                                              uint32 = 117
	AGGREGATED_URRS_TYPE                                                          uint32 = 118
	MULTIPLIER_TYPE                                                               uint32 = 119
	AGGREGATED_URR_ID_TYPE                                                        uint32 = 120
	SUBSEQUENT_VOLUME_QUOTA_TYPE                                                  uint32 = 121
	SUBSEQUENT_TIME_QUOTA_TYPE                                                    uint32 = 122
	RQI_TYPE                                                                      uint32 = 123
	QFI_TYPE                                                                      uint32 = 124
	QUERY_URR_REFERENCE_TYPE                                                      uint32 = 125
	ADDITIONAL_USAGE_REPORTS_INFORMATION_TYPE                                     uint32 = 126
	CREATE_TRAFFIC_ENDPOINT_TYPE                                                  uint32 = 127
	CREATED_TRAFFIC_ENDPOINT_TYPE                                                 uint32 = 128
	UPDATE_TRAFFIC_ENDPOINT_TYPE                                                  uint32 = 129
	REMOVE_TRAFFIC_ENDPOINT_TYPE                                                  uint32 = 130
	TRAFFIC_ENDPOINT_ID_TYPE                                                      uint32 = 131
	ETHERNET_PACKET_FILTER_TYPE                                                   uint32 = 132
	MAC_ADDRESS_TYPE                                                              uint32 = 133
	C_TAG_TYPE                                                                    uint32 = 134
	S_TAG_TYPE                                                                    uint32 = 135
	ETHERTYPE_TYPE                                                                uint32 = 136
	PROXYING_TYPE                                                                 uint32 = 137
	ETHERNET_FILTER_ID_TYPE                                                       uint32 = 138
	ETHERNET_FILTER_PROPERTIES_TYPE                                               uint32 = 139
	SUGGESTED_BUFFERING_PACKETS_COUNT_TYPE                                        uint32 = 140
	USER_ID_TYPE                                                                  uint32 = 141
	ETHERNET_PDU_SESSION_INFORMATION_TYPE                                         uint32 = 142
	ETHERNET_TRAFFIC_INFORMATION_TYPE                                             uint32 = 143
	MAC_ADDRESSES_DETECTED_TYPE                                                   uint32 = 144
	MAC_ADDRESSES_REMOVED_TYPE                                                    uint32 = 145
	ETHERNET_INACTIVITY_TIMER_TYPE                                                uint32 = 146
	ADDITIONAL_MONITORING_TIME_TYPE                                               uint32 = 147
	EVENT_QUOTA_TYPE                                                              uint32 = 148
	EVENT_THRESHOLD_TYPE                                                          uint32 = 149
	SUBSEQUENT_EVENT_QUOTA_TYPE                                                   uint32 = 150
	SUBSEQUENT_EVENT_THRESHOLD_TYPE                                               uint32 = 151
	TRACE_INFORMATION_TYPE                                                        uint32 = 152
	FRAMED_ROUTE_TYPE                                                             uint32 = 153
	FRAMED_ROUTING_TYPE                                                           uint32 = 154
	FRAMED_IPV6_ROUTE_TYPE                                                        uint32 = 155
	TIME_STAMP_TYPE                                                               uint32 = 156
	AVERAGING_WINDOW_TYPE                                                         uint32 = 157
	PAGING_POLICY_INDICATOR_TYPE                                                  uint32 = 158
	APN_DNN_TYPE                                                                  uint32 = 159
	_INTERFACE_TYPE_TYPE                                                          uint32 = 160
	PFCPSRREQ_FLAGS_TYPE                                                          uint32 = 161
	PFCPAUREQ_FLAGS_TYPE                                                          uint32 = 162
	ACTIVATION_TIME_TYPE                                                          uint32 = 163
	DEACTIVATION_TIME_TYPE                                                        uint32 = 164
	CREATE_MAR_TYPE                                                               uint32 = 165
	_ACCESS_FORWARDING_ACTION_INFORMATION_TYPE                                    uint32 = 166
	NON__ACCESS_FORWARDING_ACTION_INFORMATION_TYPE                                uint32 = 167
	REMOVE_MAR_TYPE                                                               uint32 = 168
	UPDATE_MAR_TYPE                                                               uint32 = 169
	MAR_ID_TYPE                                                                   uint32 = 170
	STEERING_FUNCTIONALITY_TYPE                                                   uint32 = 171
	STEERING_MODE_TYPE                                                            uint32 = 172
	WEIGHT_TYPE                                                                   uint32 = 173
	PRIORITY_TYPE                                                                 uint32 = 174
	UPDATE__ACCESS_FORWARDING_ACTION_INFORMATION_TYPE                             uint32 = 175
	UPDATE_NON__ACCESS_FORWARDING_ACTION_INFORMATION_TYPE                         uint32 = 176
	UE_IP_ADDRESS_POOL_IDENTITY_TYPE                                              uint32 = 177
	ALTERNATIVE_SMF_IP_ADDRESS_TYPE                                               uint32 = 178
	PACKET_REPLICATION_AND_DETECTION_CARRY_ON_INFORMATION_TYPE                    uint32 = 179
	SMF_SET_ID_TYPE                                                               uint32 = 180
	QUOTA_VALIDITY_TIME_TYPE                                                      uint32 = 181
	NUMBER_OF_REPORTS_TYPE                                                        uint32 = 182
	PFCP_SESSION_RETENTION_INFORMATION_WITHIN_PFCP_ASSOCIATION_SETUP_REQUEST_TYPE uint32 = 183
	PFCPASRSP_FLAGS_TYPE                                                          uint32 = 184
	CP_PFCP_ENTITY_IP_ADDRESS_TYPE                                                uint32 = 185
	PFCPSEREQ_FLAGS_TYPE                                                          uint32 = 186
	USER_PLANE_PATH_RECOVERY_REPORT_TYPE                                          uint32 = 187
	IP_MULTICAST_ADDRESSING_INFO_WITHIN_PFCP_SESSION_ESTABLISHMENT_REQUEST_TYPE   uint32 = 188
	JOIN_IP_MULTICAST_INFORMATION_IE_WITHIN_USAGE_REPORT_TYPE                     uint32 = 189
	LEAVE_IP_MULTICAST_INFORMATION_IE_WITHIN_USAGE_REPORT_TYPE                    uint32 = 190
	IP_MULTICAST_ADDRESS_TYPE                                                     uint32 = 191
	SOURCE_IP_ADDRESS_TYPE                                                        uint32 = 192
	PACKET_RATE_STATUS_TYPE                                                       uint32 = 193
	CREATE_BRIDGE_INFO_FOR_TSC_TYPE                                               uint32 = 194
	CREATED_BRIDGE_INFO_FOR_TSC_TYPE                                              uint32 = 195
	DS_TT_PORT_NUMBER_TYPE                                                        uint32 = 196
	NW_TT_PORT_NUMBER_TYPE                                                        uint32 = 197
	FIVEGS_USER_PLANE_NODE_TYPE                                                   uint32 = 198
	TSC_MANAGEMENT_INFORMATION_IE_WITHIN_PFCP_SESSION_MODIFICATION_REQUEST_TYPE   uint32 = 199
	TSC_MANAGEMENT_INFORMATION_IE_WITHIN_PFCP_SESSION_MODIFICATION_RESPONSE_TYPE  uint32 = 200
	TSC_MANAGEMENT_INFORMATION_IE_WITHIN_PFCP_SESSION_REPORT_REQUEST_TYPE         uint32 = 201
	PORT_MANAGEMENT_INFORMATION_CONTAINER_TYPE                                    uint32 = 202
	CLOCK_DRIFT_CONTROL_INFORMATION_TYPE                                          uint32 = 203
	REQUESTED_CLOCK_DRIFT_INFORMATION_TYPE                                        uint32 = 204
	CLOCK_DRIFT_REPORT_TYPE                                                       uint32 = 205
	TIME_DOMAIN_NUMBER_TYPE                                                       uint32 = 206
	TIME_OFFSET_THRESHOLD_TYPE                                                    uint32 = 207
	CUMULATIVE_RATERATIO_THRESHOLD_TYPE                                           uint32 = 208
	TIME_OFFSET_MEASUREMENT_TYPE                                                  uint32 = 209
	CUMULATIVE_RATERATIO_MEASUREMENT_TYPE                                         uint32 = 210
	REMOVE_SRR_TYPE                                                               uint32 = 211
	CREATE_SRR_TYPE                                                               uint32 = 212
	UPDATE_SRR_TYPE                                                               uint32 = 213
	SESSION_REPORT_TYPE                                                           uint32 = 214
	SRR_ID_TYPE                                                                   uint32 = 215
	ACCESS_AVAILABILITY_CONTROL_INFORMATION_TYPE                                  uint32 = 216
	REQUESTED_ACCESS_AVAILABILITY_INFORMATION_TYPE                                uint32 = 217
	ACCESS_AVAILABILITY_REPORT_TYPE                                               uint32 = 218
	ACCESS_AVAILABILITY_INFORMATION_TYPE                                          uint32 = 219
	PROVIDE_ATSSS_CONTROL_INFORMATION_TYPE                                        uint32 = 220
	ATSSS_CONTROL_PARAMETERS_TYPE                                                 uint32 = 221
	MPTCP_CONTROL_INFORMATION_TYPE                                                uint32 = 222
	ATSSS_LL_CONTROL_INFORMATION_TYPE                                             uint32 = 223
	PMF_CONTROL_INFORMATION_TYPE                                                  uint32 = 224
	MPTCP_PARAMETERS_TYPE                                                         uint32 = 225
	ATSSS_LL_PARAMETERS_TYPE                                                      uint32 = 226
	PMF_PARAMETERS_TYPE                                                           uint32 = 227
	MPTCP_ADDRESS_INFORMATION_TYPE                                                uint32 = 228
	UE_LINK_SPECIFIC_IP_ADDRESS_TYPE                                              uint32 = 229
	PMF_ADDRESS_INFORMATION_TYPE                                                  uint32 = 230
	ATSSS_LL_INFORMATION_TYPE                                                     uint32 = 231
	DATA_NETWORK_ACCESS_IDENTIFIER_TYPE                                           uint32 = 232
	UE_IP_ADDRESS_POOL_INFORMATION_TYPE                                           uint32 = 233
	AVERAGE_PACKET_DELAY_TYPE                                                     uint32 = 234
	MINIMUM_PACKET_DELAY_TYPE                                                     uint32 = 235
	MAXIMUM_PACKET_DELAY_TYPE                                                     uint32 = 236
	QOS_REPORT_TRIGGER_TYPE                                                       uint32 = 237
	GTP_U_PATH_QOS_CONTROL_INFORMATION_TYPE                                       uint32 = 238
	GTP_U_PATH_QOS_REPORT_PFCP_NODE_REPORT_REQUEST_TYPE                           uint32 = 239
	QOS_INFORMATION_IN_GTP_U_PATH_QOS_REPORT_TYPE                                 uint32 = 240
	GTP_U_PATH_INTERFACE_TYPE_TYPE                                                uint32 = 241
	QOS_MONITORING_PER_QOS_FLOW_CONTROL_INFORMATION_TYPE                          uint32 = 242
	REQUESTED_QOS_MONITORING_TYPE                                                 uint32 = 243
	REPORTING_FREQUENCY_TYPE                                                      uint32 = 244
	PACKET_DELAY_THRESHOLDS_TYPE                                                  uint32 = 245
	MINIMUM_WAIT_TIME_TYPE                                                        uint32 = 246
	QOS_MONITORING_REPORT_TYPE                                                    uint32 = 247
	QOS_MONITORING_MEASUREMENT_TYPE                                               uint32 = 248
	MT_EDT_CONTROL_INFORMATION_TYPE                                               uint32 = 249
	DL_DATA_PACKETS_SIZE_TYPE                                                     uint32 = 250
	QER_CONTROL_INDICATIONS_TYPE                                                  uint32 = 251
	PACKET_RATE_STATUS_REPORT_TYPE                                                uint32 = 252
	NF_INSTANCE_ID_TYPE                                                           uint32 = 253
	ETHERNET_CONTEXT_INFORMATION_TYPE                                             uint32 = 254
	REDUNDANT_TRANSMISSION_PARAMETERS_TYPE                                        uint32 = 255
	UPDATED_PDR_TYPE                                                              uint32 = 256
	S_NSSAI_TYPE                                                                  uint32 = 257
	IP_VERSION_TYPE                                                               uint32 = 258
	PFCPASREQ_FLAGS_TYPE                                                          uint32 = 259
	DATA_STATUS_TYPE                                                              uint32 = 260
	PROVIDE_RDS_CONFIGURATION_INFORMATION_TYPE                                    uint32 = 261
	RDS_CONFIGURATION_INFORMATION_TYPE                                            uint32 = 262
	QUERY_PACKET_RATE_STATUS_IE_WITHIN_PFCP_SESSION_MODIFICATION_REQUEST_TYPE     uint32 = 263
	PACKET_RATE_STATUS_REPORT_IE_WITHIN_PFCP_SESSION_MODIFICATION_RESPONSE_TYPE   uint32 = 264
	MPTCP_APPLICABLE_INDICATION_TYPE                                              uint32 = 265
	BRIDGE_MANAGEMENT_INFORMATION_CONTAINER_TYPE                                  uint32 = 266
	UE_IP_ADDRESS_USAGE_INFORMATION_TYPE                                          uint32 = 267
	NUMBER_OF_UE_IP_ADDRESSES_TYPE                                                uint32 = 268
	VALIDITY_TIMER_TYPE                                                           uint32 = 269
	REDUNDANT_TRANSMISSION_FORWARDING_PARAMETERS_TYPE                             uint32 = 270
	TRANSPORT_DELAY_REPORTING_TYPE                                                uint32 = 271
	PARTIAL_FAILURE_INFORMATION_TYPE                                              uint32 = 272
	SPARE_TYPE                                                                    uint32 = 273
	OFFENDING_IE_INFORMATION_TYPE                                                 uint32 = 274
	RAT_TYPE_TYPE                                                                 uint32 = 275
	L2TP_TUNNEL_INFORMATION_TYPE                                                  uint32 = 276
	L2TP_SESSION_INFORMATION_TYPE                                                 uint32 = 277
	L2TP_USER_AUTHENTICATION_IE_TYPE                                              uint32 = 278
	CREATED_L2TP_SESSION_TYPE                                                     uint32 = 279
	LNS_ADDRESS_TYPE                                                              uint32 = 280
	TUNNEL_PREFERENCE_TYPE                                                        uint32 = 281
	CALLING_NUMBER_TYPE                                                           uint32 = 282
	CALLED_NUMBER_TYPE                                                            uint32 = 283
	L2TP_SESSION_INDICATIONS_TYPE                                                 uint32 = 284
	DNS_SERVER_ADDRESS_TYPE                                                       uint32 = 285
	NBNS_SERVER_ADDRESS_TYPE                                                      uint32 = 286
	MAXIMUM_RECEIVE_UNIT_TYPE                                                     uint32 = 287
	THRESHOLDS_TYPE                                                               uint32 = 288
	STEERING_MODE_INDICATOR_TYPE                                                  uint32 = 289
	PFCP_SESSION_CHANGE_INFO_TYPE                                                 uint32 = 290
	GROUP_ID_TYPE                                                                 uint32 = 291
	CP_IP_ADDRESS_TYPE                                                            uint32 = 292
	IP_ADDRESS_AND_PORT_NUMBER_REPLACEMENT_TYPE                                   uint32 = 293
	DNS_QUERY_FILTER_TYPE                                                         uint32 = 294
	DIRECT_REPORTING_INFORMATION_TYPE                                             uint32 = 295
	EVENT_NOTIFICATION_URI_TYPE                                                   uint32 = 296
	NOTIFICATION_CORRELATION_ID_TYPE                                              uint32 = 297
	REPORTING_FLAGS_TYPE                                                          uint32 = 298
	PREDEFINED_RULES_NAME_TYPE                                                    uint32 = 299
	MBS_SESSION_N4MB_CONTROL_INFORMATION_TYPE                                     uint32 = 300
	MBS_MULTICAST_PARAMETERS_TYPE                                                 uint32 = 301
	ADD_MBS_UNICAST_PARAMETERS_TYPE                                               uint32 = 302
	MBS_SESSION_N4MB_INFORMATION_TYPE                                             uint32 = 303
	REMOVE_MBS_UNICAST_PARAMETERS_TYPE                                            uint32 = 304
	MBS_SESSION_IDENTIFIER_TYPE                                                   uint32 = 305
	MULTICAST_TRANSPORT_INFORMATION_TYPE                                          uint32 = 306
	MBSN4MBREQ_FLAGS_TYPE                                                         uint32 = 307
	LOCAL_INGRESS_TUNNEL_TYPE                                                     uint32 = 308
	MBS_UNICAST_PARAMETERS_ID_TYPE                                                uint32 = 309
	MBS_SESSION_N4_CONTROL_INFORMATION_TYPE                                       uint32 = 310
	MBS_SESSION_N4_INFORMATION_TYPE                                               uint32 = 311
	MBSN4RESP_FLAGS_TYPE                                                          uint32 = 312
	TUNNEL_PASSWORD_TYPE                                                          uint32 = 313
	AREA_SESSION_ID_TYPE                                                          uint32 = 314
	PEER_UP_RESTART_REPORT_TYPE                                                   uint32 = 315
	DSCP_TO_PPI_CONTROL_INFORMATION_TYPE                                          uint32 = 316
	DSCP_TO_PPI_MAPPING_INFORMATION_TYPE                                          uint32 = 317
	PFCPSDRSP_FLAGS_TYPE                                                          uint32 = 318
	QER_INDICATIONS_TYPE                                                          uint32 = 319
	VENDOR_SPECIFIC_NODE_REPORT_TYPE_TYPE                                         uint32 = 320
	CONFIGURED_TIME_DOMAIN_TYPE                                                   uint32 = 321
)
