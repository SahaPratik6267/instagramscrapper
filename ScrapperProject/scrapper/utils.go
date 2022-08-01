package scrapper

import "time"

func parseProfile(user LegacyUser) Profile {
	profile := Profile{
		Avatar:         user.ProfileImageURLHTTPS,
		Banner:         user.ProfileBannerURL,
		Biography:      user.Description,
		FollowersCount: user.FollowersCount,
		FollowingCount: user.FavouritesCount,
		FriendsCount:   user.FriendsCount,
		IsPrivate:      user.Protected,
		IsVerified:     user.Verified,
		LikesCount:     user.FavouritesCount,
		ListedCount:    user.ListedCount,
		Location:       user.Location,
		Name:           user.Name,
		PinnedTweetIDs: user.PinnedTweetIdsStr,
		TweetsCount:    user.StatusesCount,
		URL:            "https://twitter.com/" + user.ScreenName,
		UserID:         user.IDStr,
		Username:       user.ScreenName,
	}

	tm, err := time.Parse(time.RubyDate, user.CreatedAt)
	if err == nil {
		tm = tm.UTC()
		profile.Joined = &tm
	}

	if len(user.Entities.URL.Urls) > 0 {
		profile.Website = user.Entities.URL.Urls[0].ExpandedURL
	}

	return profile
}
