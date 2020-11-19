package mode

type Item struct {
	PageName string `json:"pageName"`
	Mods     struct {
		Shopcombotip struct {
			Status string `json:"status"`
		} `json:"shopcombotip"`
		Phonenav struct {
			Status string `json:"status"`
		} `json:"phonenav"`
		Debugbar struct {
			Status string `json:"status"`
		} `json:"debugbar"`
		Shopcombo struct {
			Status string `json:"status"`
		} `json:"shopcombo"`
		Itemlist struct {
			Status string `json:"status"`
			Data   struct {
				PostFeeText string `json:"postFeeText"`
				Trace       string `json:"trace"`
				Auctions    []struct {
					I2ITags struct {
						Samestyle struct {
							URL string `json:"url"`
						} `json:"samestyle"`
						Similar struct {
							URL string `json:"url"`
						} `json:"similar"`
					} `json:"i2iTags"`
					P4PTags      []interface{} `json:"p4pTags"`
					Nid          string        `json:"nid"`
					Category     string        `json:"category"`
					Pid          string        `json:"pid"`
					Title        string        `json:"title"`
					RawTitle     string        `json:"raw_title"`
					PicURL       string        `json:"pic_url"`
					DetailURL    string        `json:"detail_url"`
					ViewPrice    string        `json:"view_price"`
					ViewFee      string        `json:"view_fee"`
					ItemLoc      string        `json:"item_loc"`
					ViewSales    string        `json:"view_sales"`
					CommentCount string        `json:"comment_count"`
					UserID       string        `json:"user_id"`
					Nick         string        `json:"nick"`
					Shopcard     struct {
						LevelClasses []struct {
							LevelClass string `json:"levelClass"`
						} `json:"levelClasses"`
						IsTmall         bool   `json:"isTmall"`
						Delivery        []int  `json:"delivery"`
						Description     []int  `json:"description"`
						Service         []int  `json:"service"`
						EncryptedUserID string `json:"encryptedUserId"`
						SellerCredit    int    `json:"sellerCredit"`
						TotalRate       int    `json:"totalRate"`
					} `json:"shopcard"`
					Icon       []interface{} `json:"icon"`
					CommentURL string        `json:"comment_url"`
					ShopLink   string        `json:"shopLink"`
					Risk       string        `json:"risk"`
				} `json:"auctions"`
				RecommendAuctions []interface{} `json:"recommendAuctions"`
				IsSameStyleView   bool          `json:"isSameStyleView"`
				Sellers           []interface{} `json:"sellers"`
				Query             string        `json:"query"`
				SpmModID          string        `json:"spmModId"`
			} `json:"data"`
		} `json:"itemlist"`
		Bottomsearch struct {
			Status string `json:"status"`
			Data   struct {
				Query         string `json:"query"`
				ShowSearchBox bool   `json:"showSearchBox"`
			} `json:"data"`
		} `json:"bottomsearch"`
		Tips struct {
			Status string `json:"status"`
		} `json:"tips"`
		Feedback struct {
			Status string `json:"status"`
			Data   struct {
				Render   bool   `json:"render"`
				UseOld   bool   `json:"useOld"`
				ShowType string `json:"showType"`
				Pingce   string `json:"pingce"`
			} `json:"data"`
		} `json:"feedback"`
		Sc struct {
			Status string `json:"status"`
		} `json:"sc"`
		Navtabtags struct {
			Status string `json:"status"`
		} `json:"navtabtags"`
		Bgshopstar struct {
			Status string `json:"status"`
		} `json:"bgshopstar"`
		Spuseries struct {
			Status string `json:"status"`
		} `json:"spuseries"`
		Related struct {
			Status string `json:"status"`
			Data   struct {
				Words []struct {
					Text        string `json:"text"`
					IsHighlight bool   `json:"isHighlight"`
					Href        string `json:"href"`
				} `json:"words"`
			} `json:"data"`
		} `json:"related"`
		Tab struct {
			Status string `json:"status"`
			Data   struct {
				Tabs []struct {
					Type     string `json:"type"`
					Weight   int    `json:"weight"`
					Name     string `json:"name"`
					ID       string `json:"id"`
					Trace    string `json:"trace"`
					Href     string `json:"href"`
					Text     string `json:"text"`
					IsActive bool   `json:"isActive"`
					SpmID    string `json:"spmId"`
				} `json:"tabs"`
				SpmModID string `json:"spmModId"`
			} `json:"data"`
		} `json:"tab"`
		Pager struct {
			Status string `json:"status"`
			Data   struct {
				PageSize    int `json:"pageSize"`
				TotalPage   int `json:"totalPage"`
				CurrentPage int `json:"currentPage"`
				TotalCount  int `json:"totalCount"`
			} `json:"data"`
		} `json:"pager"`
		Apasstips struct {
			Status string `json:"status"`
		} `json:"apasstips"`
		Tbcode struct {
			Status string `json:"status"`
		} `json:"tbcode"`
		Vbaby struct {
			Status string `json:"status"`
		} `json:"vbaby"`
		Hongbao struct {
			Status string `json:"status"`
		} `json:"hongbao"`
		Nav struct {
			Status string `json:"status"`
			Data   struct {
				Common []struct {
					Text    string `json:"text"`
					Type    string `json:"type"`
					IsMulti bool   `json:"isMulti"`
					Sub     []struct {
						Text         string `json:"text"`
						IsExpandShow bool   `json:"isExpandShow"`
						Key          string `json:"key"`
						Value        string `json:"value"`
						Trace        string `json:"trace"`
						TraceData    struct {
							Click string `json:"click"`
						} `json:"traceData"`
					} `json:"sub"`
					ForceShowMore bool   `json:"forceShowMore"`
					Trace         string `json:"trace"`
				} `json:"common"`
				Breadcrumbs struct {
					Catpath []struct {
						Catid string `json:"catid"`
						Name  string `json:"name"`
					} `json:"catpath"`
				} `json:"breadcrumbs"`
			} `json:"data"`
		} `json:"nav"`
		Sortbar struct {
			Status string `json:"status"`
			Data   struct {
				SortList []struct {
					Name         string `json:"name"`
					Tip          string `json:"tip"`
					Trace        string `json:"trace"`
					IsActive     bool   `json:"isActive"`
					Value        string `json:"value"`
					Key          string `json:"key"`
					DropdownList []struct {
						Name  string `json:"name"`
						Tip   string `json:"tip"`
						Value string `json:"value"`
						Trace string `json:"trace"`
					} `json:"dropdownList,omitempty"`
				} `json:"sortList"`
				Pager struct {
					PageSize    int `json:"pageSize"`
					TotalPage   int `json:"totalPage"`
					CurrentPage int `json:"currentPage"`
					TotalCount  int `json:"totalCount"`
				} `json:"pager"`
				Price struct {
				} `json:"price"`
				Filter []struct {
					IsActive  bool   `json:"isActive"`
					Value     string `json:"value"`
					Title     string `json:"title"`
					Key       string `json:"key"`
					Trace     string `json:"trace"`
					TraceData struct {
						Filterid string `json:"filterid"`
					} `json:"traceData"`
					IsHighlight bool   `json:"isHighlight"`
					Pos         int    `json:"pos"`
					DomID       string `json:"dom_id"`
				} `json:"filter"`
				Location struct {
					Active   string        `json:"active"`
					GuessLoc string        `json:"guessLoc"`
					UsualLoc []interface{} `json:"usualLoc"`
				} `json:"location"`
				Style string `json:"style"`
			} `json:"data"`
		} `json:"sortbar"`
		D11Filterbar struct {
			Status string `json:"status"`
		} `json:"d11filterbar"`
		Personalbar struct {
			Status string `json:"status"`
		} `json:"personalbar"`
		P4P struct {
			Status string `json:"status"`
			Data   struct {
				BaobeiExtraClass string `json:"baobeiExtraClass"`
				EtaoAds          bool   `json:"etaoAds"`
				P4Pconfig        struct {
					Keyword             string `json:"keyword"`
					KeywordGBK          string `json:"keywordGBK"`
					Catid               string `json:"catid"`
					Propertyid          string `json:"propertyid"`
					IP                  string `json:"ip"`
					Loc                 string `json:"loc"`
					Gprice              string `json:"gprice"`
					Sort                string `json:"sort"`
					Sbid                string `json:"sbid"`
					Q2Cused             int    `json:"q2cused"`
					PageNum             int    `json:"pageNum"`
					P4PbottomUp         bool   `json:"p4pbottom_up"`
					B2BShow             bool   `json:"b2b_show"`
					EtaoWanke           bool   `json:"etao_wanke"`
					EtaoEffect          bool   `json:"etao_effect"`
					Offset              int    `json:"offset"`
					Refpid              string `json:"refpid"`
					Source              string `json:"source"`
					Xmatch              int    `json:"xmatch"`
					Rn                  string `json:"rn"`
					Ismall              string `json:"ismall"`
					Srp                 string `json:"srp"`
					Tags                string `json:"tags"`
					P4PBtsinfo          string `json:"p4p_btsinfo"`
					AuctionTag          string `json:"auction_tag"`
					HasSkuPic           bool   `json:"has_sku_pic"`
					FirstpageAuctionNum string `json:"firstpage_auction_num"`
					AuctionNum          struct {
						Search int `json:"search"`
					} `json:"auction_num"`
					Time       string `json:"time"`
					TianmuType string `json:"tianmu_type"`
					Refinfo    string `json:"refinfo"`
				} `json:"p4pconfig"`
				P4Pdata string `json:"p4pdata"`
			} `json:"data"`
		} `json:"p4p"`
		Choosecar struct {
			Status string `json:"status"`
		} `json:"choosecar"`
		Shopstar struct {
			Status string `json:"status"`
		} `json:"shopstar"`
		Header struct {
			Status string `json:"status"`
			Data   struct {
				Q         string `json:"q"`
				TabParams struct {
					Js           string `json:"js"`
					StatsClick   string `json:"stats_click"`
					InitiativeID string `json:"initiative_id"`
					Ie           string `json:"ie"`
				} `json:"tabParams"`
				Dropdown []struct {
					URL      string `json:"url"`
					Text     string `json:"text"`
					Type     string `json:"type"`
					IsActive bool   `json:"isActive"`
				} `json:"dropdown"`
				ImgBtn    bool   `json:"imgBtn"`
				UploadURL string `json:"uploadUrl"`
				Hb        bool   `json:"hb"`
				HbV       string `json:"hb_v"`
			} `json:"data"`
		} `json:"header"`
		Spucombo struct {
			Status string `json:"status"`
		} `json:"spucombo"`
		Supertab struct {
			Status string `json:"status"`
		} `json:"supertab"`
		Navtablink struct {
			Status string `json:"status"`
			Data   struct {
				Df struct {
					Active bool   `json:"active"`
					PbURL  string `json:"pbUrl"`
				} `json:"df"`
			} `json:"data"`
		} `json:"navtablink"`
		Noresult struct {
			Status string `json:"status"`
		} `json:"noresult"`
	} `json:"mods"`
	MainInfo struct {
		CurrentURL string `json:"currentUrl"`
		ModLinks   struct {
			Filter     string `json:"filter"`
			Default    string `json:"default"`
			Nav        string `json:"nav"`
			Breadcrumb string `json:"breadcrumb"`
			Pager      string `json:"pager"`
			Tab        string `json:"tab"`
			Sortbar    string `json:"sortbar"`
		} `json:"modLinks"`
		SrpGlobal struct {
			Q            string `json:"q"`
			EncodeQ      string `json:"encode_q"`
			Utf8Q        string `json:"utf8_q"`
			Cat          string `json:"cat"`
			S            int    `json:"s"`
			Tnk          string `json:"tnk"`
			Bucketid     int    `json:"bucketid"`
			MultiBucket  string `json:"multi_bucket"`
			Style        string `json:"style"`
			InitiativeID string `json:"initiative_id"`
			Machine      string `json:"machine"`
			Buckets      string `json:"buckets"`
			SpURL        string `json:"sp_url"`
			SrpName      string `json:"srpName"`
		} `json:"srpGlobal"`
		TraceInfo struct {
			PvStat    string `json:"pvStat"`
			TraceData struct {
				Catdirect           string        `json:"catdirect"`
				Remoteip            string        `json:"remoteip"`
				RewriteStatus       string        `json:"rewriteStatus"`
				TabType             string        `json:"tabType"`
				IsRs                string        `json:"is_rs"`
				CatpredictBury      string        `json:"catpredict_bury"`
				Hostname            string        `json:"hostname"`
				FewWordsPv          string        `json:"fewWords_pv"`
				ActivityClick       []string      `json:"activityClick"`
				LastAlitrackid      string        `json:"lastAlitrackid"`
				AtLflog             string        `json:"at_lflog"`
				ListModel           string        `json:"list_model"`
				PageSize            string        `json:"page_size"`
				RsPositions         []string      `json:"rsPositions"`
				IfTank              string        `json:"if_tank"`
				Rsshop              string        `json:"rsshop"`
				Alitrackid          string        `json:"alitrackid"`
				Query               string        `json:"query"`
				PriceRank           string        `json:"price_rank"`
				Sort                string        `json:"sort"`
				CatLevelOne         string        `json:"catLevelOne"`
				AuctionNids         []string      `json:"auctionNids"`
				IfDoufuAuction      []string      `json:"ifDoufuAuction"`
				AtHost              string        `json:"at_host"`
				QuerytypeBury       string        `json:"querytype_bury"`
				AllOldBiz30Day      []string      `json:"allOldBiz30Day"`
				TdTags              string        `json:"tdTags"`
				RelateHotTrace      []string      `json:"relateHotTrace"`
				TotalHits           string        `json:"totalHits"`
				AllCategories       []string      `json:"allCategories"`
				AuctionIconServices []string      `json:"auctionIconServices"`
				RewriteQuery        string        `json:"rewriteQuery"`
				NavPopup            string        `json:"navPopup"`
				Rn                  string        `json:"rn"`
				Isp4P               []string      `json:"isp4p"`
				Rs                  string        `json:"rs"`
				NavCategory         string        `json:"navCategory"`
				Colo                string        `json:"colo"`
				AllPrices           []string      `json:"allPrices"`
				ShowCompass         string        `json:"show_compass"`
				AuctionPrices       []string      `json:"auctionPrices"`
				AuctionReturnNum    string        `json:"auctionReturnNum"`
				Multivariate        string        `json:"multivariate"`
				P4PDelTraceInfo     []interface{} `json:"p4pDelTraceInfo"`
				BucketID            string        `json:"bucketId"`
				RewriteBury         string        `json:"rewrite_bury"`
				NavEntries          string        `json:"navEntries"`
				Nick                string        `json:"nick"`
				AllPersonalUpReason []string      `json:"allPersonalUpReason"`
				AllDoufuNids        []interface{} `json:"allDoufuNids"`
				FewWordsRw          string        `json:"fewWords_rw"`
				PriceSorts          []string      `json:"priceSorts"`
				AtBucketid          string        `json:"at_bucketid"`
				Srppage             string        `json:"srppage"`
				IfRs                string        `json:"if_rs"`
				AllNids             []string      `json:"allNids"`
				Cat                 string        `json:"cat"`
				SpURL               string        `json:"spUrl"`
				Sort2               string        `json:"sort2"`
				QpBury              string        `json:"qp_bury"`
				DoufuAuctionNum     string        `json:"doufuAuctionNum"`
				AtColo              string        `json:"at_colo"`
				Bandit              string        `json:"bandit"`
				RsCount             string        `json:"rs_count"`
				NavType             string        `json:"navType"`
				HasSkuPic           string        `json:"has_sku_pic"`
				FromPos             string        `json:"from_pos"`
				StatsClick          string        `json:"statsClick"`
				FewWords            []string      `json:"fewWords"`
				NavHasRanked        string        `json:"navHasRanked"`
				AllDoufuPrices      []interface{} `json:"allDoufuPrices"`
				RsKeywords          []string      `json:"rsKeywords"`
				TagList             []string      `json:"tagList"`
				AuctionNicks        []string      `json:"auctionNicks"`
				SpSellerTypes       []string      `json:"sp_seller_types"`
				CatdirectForMaidian string        `json:"catdirectForMaidian"`
				Qinfo               string        `json:"qinfo"`
				NoResultCode        string        `json:"noResultCode"`
				Apass               string        `json:"apass"`
				SpuCombo            string        `json:"spu_combo"`
				AllTags             []string      `json:"allTags"`
				MultiBucket         string        `json:"multi_bucket"`
				NavStatus           string        `json:"navStatus"`
			} `json:"traceData"`
		} `json:"traceInfo"`
		RemainMods []interface{} `json:"remainMods"`
	} `json:"mainInfo"`
	Feature struct {
		WebpOff     bool `json:"webpOff"`
		RetinaOff   bool `json:"retinaOff"`
		ShopcardOff bool `json:"shopcardOff"`
	} `json:"feature"`
}