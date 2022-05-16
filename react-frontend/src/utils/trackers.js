import analyticsService from '@razorpay/universe-utils/analytics';
import { getUserId } from './helper';

export const initLumberjack = () => {
	analyticsService.init({
		lumberjackAppName: 'external_status_page',
		lumberjackApiKey: window.LUMBERJACK_API_KEY,
		lumberjackApiUrl: window.LUMBERJACK_API_URL,
	});
};

export const analyticsTrack = (trackObj) => {
	analyticsService.track({
		...trackObj,
		properties: {
			...trackObj.properties,
			userId: getUserId(),
		},
	});
};
