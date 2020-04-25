CREATE TABLE reports(
	id UUID PRIMARY KEY,
	source TEXT NOT NULL,
	source_identity_id UUID NOT NULL,
	reference JSONB NOT NULL,
    state TEXT NOT NULL,
    payload JSONB NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP
);

INSERT INTO reports (
    id,
    source,
    source_identity_id,
    reference,
    state,
    payload,
    created_at,
    updated_at
) VALUES (
    '0103e005-b762-485f-8f7e-722019d4f302',
    'REPORT',
    '6750b4d5-4cb5-45f0-8b60-61be2072cce2',
    '{
        "reference_id": "6706b3ba-bf36-4ad4-9b9d-4ebf4f4e2429",
        "reference_type": "REPORT"
    }',
    'OPEN',
    '{
        "source": "REPORT",
        "report_type": "SPAM",
        "message": null,
        "report_id": "6706b3ba-bf36-4ad4-9b9d-4ebf4f4e2429",
        "reference_resource_id": "a03411ce-0197-49a2-86d4-55e06aa52e79",
        "reference_resource_type": "REPLY"
    }',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
),
(
    '0595df57-e5e2-4acd-8fd7-a234fa70ab14',
    'REPORT',
    '6750b4d5-4cb5-45f0-8b60-61be2072cce2',
    '{
        "referenceId": "d9865001-c46e-4fd5-b810-4310fa41cf3d",
        "referenceType": "REPORT"
    }',
    'OPEN',
    '{
        "source": "REPORT",
        "report_type": "VIOLATES_POLICIES",
        "message": "hjasds asjdkjsa daskds dasjkds",
        "report_id": "43395ec2-2895-4f66-8f96-26d2a411cae8",
        "reference_resource_id": "a9512d24-6240-4da4-b792-60ddc59b452e",
        "reference_resource_type": "ARTICLE"
    }',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
),
(
    '0bf4a85b-6da3-4eab-a740-55ec51006d0e',
    'REPORT',
    '6750b4d5-4cb5-45f0-8b60-61be2072cce2',
    '{
        "reference_id": "d9865001-c46e-4fd5-b810-4310fa41cf3d",
        "reference_type": "REPORT"
    }',
    'OPEN',
    '{
        "source": "REPORT",
        "report_type": "VIOLATES_POLICIES",
        "message": "mesage message",
        "report_id": "233d9a03-d68f-44ee-9977-bf9b4269c428",
        "reference_resource_id": "a9512d24-6240-4da4-b792-60ddc59b452e",
        "reference_resource_type": "ARTICLE"
    }',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
),
(
    '015bfeed-34a5-492d-bf4e-51a9afffe1ea',
    'REPORT',
    '4bd630eb-4b36-4038-aa8e-e58c4025de1f',
    '{
        "reference_id": "7274d582-9a1e-42bd-aa0f-f563904bfbab",
        "reference_type": "REPORT"
    }',
    'OPEN',
    '{
        "source": "REPORT",
        "report_type": "INFRINGES_PROPERTY",
        "message": "mesage message",
        "report_id": "7274d582-9a1e-42bd-aa0f-f563904bfbab",
        "reference_resource_id": "1573590e-f4bc-4cf4-82af-378a83fea5ab",
        "reference_resource_type": "POST"
    }',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
),
(
    '030015d0-097c-4991-892d-06aff536bb6c',
    'REPORT',
    'd0ba4c4a-39da-4d2c-8934-80652da104fe',
    '{
        "reference_id": "7b4cf776-49d5-4c24-8837-153fc34d6d2c",
        "reference_type": "REPORT"
    }',
    'OPEN',
    '{
        "source": "REPORT",
        "report_type": "SPAM",
        "message": "important",
        "report_id": "7b4cf776-49d5-4c24-8837-153fc34d6d2c",
        "reference_resource_id": "60c858e0-6961-4d20-a949-ef3382b2e8e4",
        "reference_resource_type": "POST"
    }',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
),
(
    '033bedc5-c77f-4069-b162-7a6cf7b835b2',
    'REPORT',
    '6750b4d5-4cb5-45f0-8b60-61be2072cce2',
    '{
        "reference_id": "d9865001-c46e-4fd5-b810-4310fa41cf3d",
        "reference_type": "REPORT"
    }',
    'OPEN',
    '{
        "source": "REPORT",
        "report_type": "VIOLATES_POLICIES",
        "message": "very important",
        "report_id": "d9865001-c46e-4fd5-b810-4310fa41cf3d",
        "reference_resource_id": "c8ba7a4b-1aa6-471b-bf91-748f915af1fb",
        "reference_resource_type": "POST"
    }',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
),
(
    '04c10147-fe29-435b-8973-94da0c228f74',
    'REPORT',
    '6750b4d5-4cb5-45f0-8b60-61be2072cce2',
    '{
        "reference_id": "65156553-b340-413d-9e30-357fa00b14a5",
        "reference_type": "REPORT"
    }',
    'OPEN',
    '{
        "source": "REPORT",
        "report_type": "INFRINGES_PROPERTY",
        "message": "important",
        "report_id": "65156553-b340-413d-9e30-357fa00b14a5",
        "reference_resource_id": "aa875f29-84dd-438d-8682-47412c8ec7d9",
        "reference_resource_type": "POST"
    }',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
),
(
    '0595df57-e5e2-4acd-8fd7-a234fa70ab18',
    'REPORT',
    '6750b4d5-4cb5-45f0-8b60-61be2072cce2',
    '{
        "reference_id": "df8168eb-5b9a-471f-bcb1-23644bbbb1b6",
        "reference_type": "REPORT"
    }',
    'OPEN',
    '{
        "source": "REPORT",
        "report_type": "VIOLATES_POLICIES",
        "message": "very important",
        "report_id": "df8168eb-5b9a-471f-bcb1-23644bbbb1b6",
        "reference_resource_id": "df8168eb-5b9a-471f-bcb1-23644bbbb1b6",
        "reference_resource_type": "POST"
    }',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
);