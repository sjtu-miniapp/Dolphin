import axios from 'axios';
import { DataType } from './types';

const BASE_URL = 'http://121.199.33.44:9999/api/v1';

export interface AddProps {
  targetOpenID: string;
  opsOpenID: string;
  opsSessionID: string;
  type: DataType;
  id: number;
}

export const addTo = async (params: AddProps): Promise<void> => {
  const { targetOpenID, opsOpenID, opsSessionID, type, id } = params;
  const suffix = type === 'group' ? 'member' : 'worker';
  const url = `${BASE_URL}/${type}/${id}/${suffix}?openid=${opsOpenID}&sid=${opsSessionID}`;
  console.log('url', url);
  const data = {
    action: 0,
    [`${suffix}s`]: [targetOpenID],
  };
  console.log('data', data);

  const response = await axios.put(url, data);
  console.log(response.data, response.status);
};
