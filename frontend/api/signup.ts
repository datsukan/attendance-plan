import axios from 'axios';

import { UnknownErrorMessage } from './error';

export const signup = async (email: string): Promise<string> => {
  const param = { email };

  try {
    await axios.post(`${process.env.NEXT_PUBLIC_API_BASE_URL}/signup`, param, {
      headers: {
        'Content-Type': 'application/json',
      },
    });

    return '';
  } catch (e) {
    if (!axios.isAxiosError(e)) {
      return UnknownErrorMessage;
    }

    const body = e.response?.data;
    if (!body) {
      return UnknownErrorMessage;
    }

    return body.message || UnknownErrorMessage;
  }
};
