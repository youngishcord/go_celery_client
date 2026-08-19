from enum import nonmember


class GoClientError(Exception):
    def __init__(self, message, code=None, details=None):
        super(GoClientError, self).__init__(message)
        self.message = message
        self.code = code
        self.details = details or {}

    def to_dict(self):
        return {
            'message': self.message,
            'code': self.code,
            'details': self.details
        }

    @classmethod
    def from_dict(cls, data):
        return cls(data['message'], data.get('code'), data.get('details'))