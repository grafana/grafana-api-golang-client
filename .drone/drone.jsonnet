local image = 'grafana/build-container:1.2.27';

local pipeline(name, trigger) = {
  kind: 'pipeline',
  type: 'docker',
  name: name,
  platform: {
    os: 'linux',
    arch: 'amd64',
  },
  trigger: trigger,
  steps: [
    {
      name: 'test',
      image: image,
      commands: [
        'make test',
      ],
    },
  ],
};

[
  pipeline('test-pr', {
    event: [
      'pull_request',
    ],
  }),

  pipeline('test-master', {
    branch: 'master',
    event: [
      'push',
    ],
  }),
]
