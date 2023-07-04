package I18n

var I18nDict_zh_HK map[uint16]string = map[uint16]string{
	ACME_FailedToGetCommand:             "未能讀取ACME指令",
	ACME_FailedToSeek:                   "無效的ACME文檔，因為seek操作失敗了。",
	ACME_StructureErrorNotice:           "文檔結構錯誤",
	ACME_UnknownCommand:                 "未知ACME指令（文檔錯誤）",
	Auth_BackendError:                   "後端錯誤",
	Auth_FailedToRequestEntry:           "未能連接伺服器入口，請檢查伺服器等級設定是否关闭及伺服器密码是否正确。",
	Auth_HelperNotCreated:               "辅助用戶尚未创建，請前往用戶中心进行创建。",
	Auth_InvalidFBVersion:               "FastBuilder 版本無效，請更新。",
	Auth_InvalidHelperUsername:          "辅助用戶的用戶名無效，請前往用戶中心进行設定。",
	Auth_InvalidToken:                   "無效Token，請重新登入。",
	Auth_InvalidUser:                    "無效帳戶，請重新登入。",
	Auth_ServerNotFound:                 "沒有偵測到伺服器，請檢查伺服器是否对所有人可見。",
	Auth_UnauthorizedRentalServerNumber: "對應伺服器尚未授權，請前往用戶中心授權。",
	Auth_UserCombined:                   "此用戶已經合併到另一個帳戶中，請使用新帳戶登入。",
	Auth_FailedToRequestEntry_TryAgain:  "未能請求伺服器入口，請稍後再試。",
	BDump_Author:                        "著者",
	BDump_EarlyEOFRightWhenOpening:      "未能讀取文檔，因為文檔过早结束，可能已經毀損。",
	BDump_FailedToGetCmd1:               "無法讀取 cmd[pos:0] 的任何參數值，文檔可能已經毀損",
	BDump_FailedToGetCmd2:               "無法讀取 cmd[pos1] 的任何參數值，文檔可能已經毀損",
	BDump_FailedToGetCmd4:               "無法讀取 cmd[pos2] 的任何參數值，文檔可能已經毀損",
	BDump_FailedToGetCmd6:               "無法讀取 cmd[pos3] 的任何參數值，文檔可能已經毀損",
	BDump_FailedToGetCmd7_0:             "無法讀取 cmd[pos4] 的任何參數值，文檔可能已經毀損",
	BDump_FailedToGetCmd7_1:             "無法讀取 cmd[pos5] 的任何參數值，文檔可能已經毀損",
	BDump_FailedToGetCmd10:              "無法讀取 cmd[pos6] 的任何參數值，文檔可能已經毀損",
	BDump_FailedToGetCmd11:              "無法讀取 cmd[pos7] 的任何參數值，文檔可能已經毀損",
	BDump_FailedToGetCmd12:              "無法讀取 cmd[pos8] 的任何參數值，文檔可能已經毀損",
	BDump_FailedToGetConstructCmd:       "無法讀取構建指令，文檔可能已經毀損",
	BDump_FailedToReadAuthorInfo:        "未能讀取著者信息，文檔可能已毀損",
	BDump_FileNotSigned:                 "文檔未簽名",
	BDump_FileSigned:                    "文檔已簽名，持有者：%s",
	BDump_NotBDX_Invheader:              "不是bdx文檔（無效header）",
	BDump_NotBDX_Invinnerheader:         "不是bdx文檔（無效innerheader）",
	BDump_SignedVerifying:               "文檔已簽名，正在驗證...",
	BDump_VerificationFailedFor:         "因 %v 未能驗證文檔簽名。",
	BDump_Warn_Reserved:                 "警告：BDump/Import：使用了保留字\n",
	CommandNotFound:                     "找不到此指令",
	ConnectionEstablished:               "成功連線到伺服器。",
	Copyright_Notice_Bouldev:            "版權所有 (c) FastBuilder DevGroup, Bouldev 2022",
	Copyright_Notice_Contrib:            "Contributors: Ruphane, CAIMEO, CMA2401PT",
	Crashed_No_Connection:               "較長時間未能建立連線",
	Crashed_OS_Windows:                  "按ENTER來退出應用程式。",
	Crashed_StackDump_And_Error:         "Stack dump 於上方顯示。錯誤为：",
	Crashed_Tip:                         "FastBuilder Phoenix 運作過程遇到問題",
	CurrentDefaultDelayMode:             "當前默認延遲模式",
	CurrentTasks:                        "任務列表：",
	DelayModeSet:                        "延遲模式已設定",
	DelayModeSet_DelayAuto:              "延遲數值已自動設定为: %d",
	DelayModeSet_ThresholdAuto:          "延遲臨界值已自動設定为: %d",
	DelaySet:                            "延遲已設定",
	DelaySetUnavailableUnderNoneMode:    "[delay set] 於 none 模式下不可用",
	DelayThreshold_OnlyDiscrete:         "延遲臨界值只可在 discrete 模式下被設定。",
	DelayThreshold_Set:                  "延遲臨界值已設定为 %d",
	ERRORStr:                            "錯誤",
	EnterPasswordForFBUC:                "請輸入你的FastBuilder用戶中心登入密码(不會顯示): ",
	Enter_FBUC_Username:                 "輸入你的FastBuilder用戶中心帳戶名: ",
	Enter_Rental_Server_Code:            "請輸入伺服器号: ",
	Enter_Rental_Server_Password:        "輸入伺服器密码 (如果沒有設定则直接按[Enter], 輸入不會顯示): ",
	ErrorIgnored:                        "已忽略此錯誤",
	Error_MapY_Exceed:                   "使用3D Map功能時，MapY 應當為 20-255 的數值（您的鍵入為 %v)",
	FBUC_LoginFailed:                    "FastBuilder用戶中心的帳戶名或密码無效",
	FBUC_Token_ErrOnCreate:              "建立Token文檔时出錯：",
	FBUC_Token_ErrOnGen:                 "產生臨時Token时出錯",
	FBUC_Token_ErrOnRemove:              "未能刪除token文檔: %v",
	FBUC_Token_ErrOnSave:                "儲存Token时出錯：",
	FileCorruptedError:                  "文檔已毀損",
	Get_Warning:                         "get指令已過時，未來將移除此指令，請遷移至set",
	IgnoredStr:                          "已忽略",
	InvalidFileError:                    "無效文檔",
	InvalidPosition:                     "未獲取到有效座標。（可忽略此錯誤）",
	Lang_Config_ErrOnCreate:             "建立語言配置文檔時出錯：%v",
	Lang_Config_ErrOnSave:               "儲存语言配置時出錯：%v",
	LanguageName:                        "繁體中文",
	LanguageUpdated:                     "语言偏好已更新",
	Logout_Done:                         "已从FastBuilder用戶中心退出登入。",
	Menu_BackButton:                     "< 返回",
	Menu_Cancel:                         "取消",
	Menu_CurrentPath:                    "當前路徑",
	Menu_ExcludeCommandsOption:          "排除指令方塊內容",
	Menu_GetEndPos:                      "獲取終點座標",
	Menu_GetPos:                         "獲取座標",
	Menu_InvalidateCommandsOption:       "指令無效化",
	Menu_Quit:                           "退出應用程式",
	Menu_StrictModeOption:               "嚴格模式",
	NotAnACMEFile:                       "所提供的文檔不是ACME結構文檔",
	Notice_CheckUpdate:                  "正在檢查更新，這可能需要一些時間…",
	Notice_iSH_Location_Service:         "您正在使用iSH模擬器，位置權限需要用於維持後台運作，除此之外沒有任何定位數據被紀錄或使用，您可以隨時關閉它。",
	Notice_OK:                           "完成\n",
	Notice_UpdateAvailable:              "有新的PhoenixBuilder版本（%s）可用。\n",
	Notice_UpdateNotice:                 "請更新本應用程式。\n",
	Notice_ZLIB_CVE:                     "您的zlib版本（%s）包含已被證實的嚴重漏洞，我們建議您更新它，避免發生意外",
	Notify_NeedOp:                       "需要 OP 權限以正常運作。",
	Notify_TurnOnCmdFeedBack:            "需要 sendcommandfeedback 为 true，我們已經為你打開該選項，使用完后請按需關閉",
	Omega_Enabled:                       "Omega系統已啟用！",
	Parsing_UnterminatedEscape:          "換碼未終止",
	Parsing_UnterminatedQuotedString:    "字符串引號部分未終止",
	PositionGot:                         "已獲得到起點座標",
	PositionGot_End:                     "已獲得終點座標",
	PositionSet:                         "已設定座標",
	PositionSet_End:                     "已設定終點座標",
	QuitCorrectly:                       "正常退出",
	Sch_FailedToResolve:                 "未能解析文檔",
	SelectLanguageOnConsole:             "請在控制台中選擇新语言",
	ServerCodeTrans:                     "伺服器號碼",
	SimpleParser_Int_ParsingFailed:      "解析器：未能處理整數型參數",
	SimpleParser_InvEnum:                "解析器：無效枚舉值，可用值有：%s.",
	SimpleParser_Invalid_decider:        "解析器：無效決定子",
	SimpleParser_Too_few_args:           "解析器：參數過少",
	Special_Startup:                     "已啟用语言：繁體中文\n",
	SysError_EACCES:                     "權限拒絕，請檢查是否已經允許該程式訪問對應文檔。",
	SysError_EBUSY:                      "文檔被佔用，請稍後再試",
	SysError_EINVAL:                     "無效文檔輸入。",
	SysError_EISDIR:                     "輸入文檔为目錄，無效輸入。",
	SysError_ENOENT:                     "對應文檔不存在。",
	SysError_ETXTBSY:                    "文檔被占用，請稍后再试。",
	SysError_HasTranslation:             "对於 %s 的文檔操作出錯：%s",
	TaskCreated:                         "任務已創建",
	TaskDisplayModeSet:                  "任務狀態顯示模式已經設定为: %s.",
	TaskFailedToParseCommand:            "未能解析指令: %v",
	TaskNotFoundMessage:                 "未能根據所提供的任務ID找到有效任務。",
	TaskPausedNotice:                    "[任務 %d] - 已暫停",
	TaskResumedNotice:                   "[任務 %d] - 已恢復",
	TaskStateLine:                       "ID %d - 命令行 \"%s\", 狀態: %s, 延遲值: %d, 延遲模式: %s, 延遲臨界值: %d",
	TaskStoppedNotice:                   "[任務 %d] - 已停止",
	TaskTTeIuKoto:                       "任務",
	TaskTotalCount:                      "總數：%d",
	TaskTypeCalculating:                 "正在計算",
	TaskTypeDied:                        "已死亡",
	TaskTypePaused:                      "已暫停",
	TaskTypeRunning:                     "運行中",
	TaskTypeSpecialTaskBreaking:         "特殊任務:正在終止",
	TaskTypeSwitchedTo:                  "任務創建類型已經切換为：%s.",
	TaskTypeUnknown:                     "未知",
	Task_D_NothingGenerated:             "[任務 %d] 沒有任何結構成功生成。",
	Task_DelaySet:                       "[任務 %d] - 延遲已設定: %d",
	Task_ResumeBuildFrom:                "從第 %v 個方塊处恢復構建",
	Task_SetDelay_Unavailable:           "[setdelay] 在 none 延遲模式下不可用",
	Task_Summary_1:                      "[任務 %d] %v 个方塊被更改。",
	Task_Summary_2:                      "[任務 %d] 用時: %v 秒",
	Task_Summary_3:                      "[任務 %d] 平均速度: %v 方塊/秒",
	UnsupportedACMEVersion:              "不支援該版本ACME結構（仅支援acme 1.2文檔版本）",
	Warning_ACME_Deprecated:             "警告 - `acme' 功能已廢棄且已移除，請遷移到BDX格式。詳情請參閱 https://github.com/LNSSPsd/PhoenixBuilder/issues/313",
	Warning_Schem_Deprecated:            "警告 - `schem' 功能已廢棄且已移除，請遷移到BDX格式。詳情請參閱 https://github.com/LNSSPsd/PhoenixBuilder/issues/313",
	Warning_UserHomeDir:                 "警告 - 無法獲取當前user主目錄，將設定homedir=\".\";\n",
}
