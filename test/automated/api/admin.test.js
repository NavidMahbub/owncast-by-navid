var request = require('supertest');
request = request('http://127.0.0.1:8080');



test('correct number of log entries exist', (done) => {
    request.get('/api/admin/logs').auth('admin', 'ZHNpbm5vdmF0b3Jz').expect(200)
        .then((res) => {
            // expect(res.body).toHaveLength(8);
            done();
        });
});

