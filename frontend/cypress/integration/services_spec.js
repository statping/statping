/// <reference types="cypress" />

import "../support/commands"

context('Services Tests', () => {

  beforeEach(() => {
    cy.restoreLocalStorageCache();
  });

  afterEach(() => {
    cy.saveLocalStorageCache();
  });

  it('should Login', () => {
    cy.visit('/login')
    cy.get('#username').clear().type('admin')
    cy.get('#password').clear().type('admin')
    cy.get('button[type="submit"]').click()

    cy.get('.navbar-brand').should('contain', 'Statping')
    cy.getCookies()

    cy.getCookies().should('have.length', 1)
  })

  it('should goto services', () => {
    cy.visit('/dashboard/services')
    cy.get(':nth-child(1) > .card-body > .table > tbody > tr').should('have.length', 5)
    cy.get('.sortable_groups > tr').should('have.length', 3)
  })

  it('should create new HTTP service', () => {
    cy.visit('/dashboard/create_service')
    cy.get(':nth-child(1) > .card-body > :nth-child(1) > .col-sm-8 > .form-control').clear().type('HTTP Service')
    cy.get('#service_type').select('http')
    cy.get('#service_url').clear().type('http://localhost:8888')
    cy.get('#service_response_code').clear().type('200')
    cy.get('#service_interval').clear().type('30')
    cy.get(':nth-child(3) > .card-body > :nth-child(2) > .col-sm-8 > .form-control').clear().type('5')
    cy.get('#permalink').clear().type('http_service')

    cy.get('#notify_after').clear().type('3')

    cy.get('button[type="submit"]').click()
  })

  it('should create new TCP service', () => {
    cy.visit('/dashboard/create_service')
    cy.get(':nth-child(1) > .card-body > :nth-child(1) > .col-sm-8 > .form-control').clear().type('TCP Service')
    cy.get('#service_type').select('tcp')
    cy.get('#service_url').clear().type('localhost')
    cy.get('#service_port').clear().type('8888')

    cy.get('#service_interval').clear().type('30')
    cy.get(':nth-child(3) > .card-body > :nth-child(2) > .col-sm-8 > .form-control').clear().type('5')
    cy.get('#notify_after').clear().type('3')

    cy.get('#permalink').clear().type('tcp_service')

    cy.get('button[type="submit"]').click()
  })

  it('should create new UDP service', () => {
    cy.visit('/dashboard/create_service')
    cy.get(':nth-child(1) > .card-body > :nth-child(1) > .col-sm-8 > .form-control').clear().type('UDP Service')
    cy.get('#service_type').select('udp')
    cy.get('#service_url').clear().type('8.8.8.8')
    cy.get('#service_port').clear().type('53')

    cy.get('#service_interval').clear().type('30')
    cy.get(':nth-child(3) > .card-body > :nth-child(2) > .col-sm-8 > .form-control').clear().type('5')
    cy.get('#notify_after').clear().type('3')

    cy.get('#permalink').clear().type('udp_service')

    cy.get('button[type="submit"]').click()
  })

  it('should create new ICMP service', () => {
    cy.visit('/dashboard/create_service')
    cy.get(':nth-child(1) > .card-body > :nth-child(1) > .col-sm-8 > .form-control').clear().type('ICMP Service')
    cy.get('#service_type').select('icmp')
    cy.get('#service_url').clear().type('8.8.8.8')

    cy.get('#service_interval').clear().type('30')
    cy.get('#notify_after').clear().type('3')

    cy.get('#permalink').clear().type('icmp_service')

    cy.get('button[type="submit"]').click()
  })

  it('should confirm new services', () => {
    cy.visit('/dashboard/services')
    cy.get(':nth-child(1) > .card-body > .table > tbody > tr').should('have.length', 9)
  })

})
