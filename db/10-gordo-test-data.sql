INSERT INTO endpoint(ep, d, base) VALUES 
	(
		'EP',
		'D',
		'BASE'
	);

INSERT INTO resource(path, ct, rt, if, anchor, endpoint) VALUES 
	(
		'coap://local-proxy-old.example.com:5683/sensors/temp',
		41,
		'temperature',
		'',
		'coap://spurious.example.com:5683',
		1
	),
	(
		'coap://local-proxy-old.example.com:5683/sensors/light',
		41,
		'light-lux',
		'sensor',
		'coap://spurious.example.com:5683',
		1
	);
