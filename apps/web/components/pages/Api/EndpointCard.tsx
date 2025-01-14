import type { SchemaObject } from 'openapi3-ts'
import type { Languages, ServiceExamples } from '@/types'
import { FC } from 'react'
import { splitEndpointTitle } from '@/utils/api'
import { RequestBlock, ResponseBlock, Card } from '@/components/ui'
import { PropertiesTable } from './PropertiesTable'

interface Props {
  apiName: string
  apiMethod: string
  apiVersion: string
  endpointName: string
  examples: ServiceExamples
  language: Languages
  onLanguageChange: (language: Languages) => void
  price: string
  requestSchema?: SchemaObject
  responseSchema?: SchemaObject
  responsePayload: Record<string, unknown>
  title: string
}

export const EndpointCard: FC<Props> = ({
  apiName,
  apiMethod,
  apiVersion,
  endpointName,
  examples,
  language,
  onLanguageChange,
  price,
  requestSchema,
  responseSchema,
  responsePayload,
  title,
}) => {
  return (
    <div className="endpoint-card" id={title.replace(/ /g, '')}>
      <div className="md:grid md:grid-cols-2">
        <div className="pb-4">
          <p className="mr-6 mt-4 float-right text-xs text-indigo-400">
            {price === 'Free' ? '0.000001' : price} credits per request
          </p>
          <h3 className="font-medium text-3xl mb-2">
            {splitEndpointTitle(title)}
          </h3>
          <h4 className="mb-4 inline-block text-xs">
            <span className="font-bold text-indigo-400">{apiMethod}</span>{' '}
            <span className="text-sm">
              /{apiVersion}/{apiName}/{endpointName}
            </span>
          </h4>
          {requestSchema && (
            <p className="mb-10  max-w-md">
              {requestSchema.description}
            </p>
          )}
          <PropertiesTable
            title="Request"
            properties={requestSchema?.properties}
          />
          <PropertiesTable
            title="Response"
            properties={responseSchema?.properties}
          />
        </div>
        <div className="pb-6">
          {requestSchema && (
            <RequestBlock
              examples={examples}
              requestSchema={requestSchema}
              apiName={apiName}
              onLanguageChange={onLanguageChange}
              language={language}
            />
          )}
          {responseSchema && (
            <ResponseBlock code={JSON.stringify(responsePayload, null, 4)} />
          )}
        </div>
      </div>
    </div>
  )
}
